package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkspaceCard(c *gin.Context) {
	var req = structs.WorkspaceCreateCardRequest{}
	var workspace_id = c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.First(&models.Workspace{}, workspace_id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_card_status := models.WorkspaceCardStatus{}

	if err := database.DB.Where("id = ? AND workspace_id = ? AND is_archived = ?", req.StatusId, workspace_id, false).First(&workspace_card_status).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Status card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		result := trx.Model(&models.WorkspaceCard{}).Where("status_id = ?", req.StatusId)

		var card = models.WorkspaceCard{
			WorkspaceId:   uint(parsed_workspace_id),
			StatusId:      uint(req.StatusId),
			UserCreatedId: uint(user_id),
			Title:         req.Title,
			Description:   req.Description,
			Order:         int(result.RowsAffected),
		}

		if err := trx.Create(&card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_card_activity := models.WorkspaceCardActivity{
			CardId:  card.Id,
			UserId:  uint(user_id),
			Type:    "general",
			Message: fmt.Sprintf("User %s created this card", user.Username),
		}

		if err := trx.Create(&workspace_card_activity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "create-workspace-card",
			Description: fmt.Sprintf("User %s has added new card to status %s", user.Username, workspace_card_status.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusCreated, structs.SuccessResponse{
			Success: true,
			Message: "Workspace card created successfully",
			Data: structs.WorkspaceCardResponse{
				Id:          card.Id,
				Title:       card.Title,
				Description: card.Description,
				Members:     []structs.Member{},
			},
		})

		return nil
	})
}
