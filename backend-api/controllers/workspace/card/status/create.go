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

func CreateWorkspaceCardStatus(c *gin.Context) {
	req := structs.WorkspaceCreateCardStatusRequest{}
	workspace_id := c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)
	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_card_status := models.WorkspaceCardStatus{
		WorkspaceId: uint(parsed_workspace_id),
		Name:        req.Name,
		Color:       req.Color,
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		list_workspace_card_status := []models.WorkspaceCardStatus{}
		trx.Find(&list_workspace_card_status, "workspace_id = ?", workspace_id)
		workspace_card_status.Order = len(list_workspace_card_status)

		if err := trx.Create(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			UserRole:    "manager",
			Action:      "create-workspace-card-status",
			Description: fmt.Sprintf("User %s has created new status", user.Username),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_log := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Success to create workspace status",
			Data: structs.WorkspaceCardStatus{
				Id:         workspace_card_status.Id,
				Order:      workspace_card_status.Order,
				Name:       workspace_card_status.Name,
				Color:      workspace_card_status.Color,
				IsArchived: workspace_card_status.IsArchived,
			},
		})

		return nil
	})
}
