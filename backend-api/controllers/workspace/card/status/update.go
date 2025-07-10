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

func UpdateWorkspaceCardStatus(c *gin.Context) {
	req := structs.WorkspaceUpdateCardStatusRequest{}
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	status_id := c.Param("status_id")

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

	workspace_card_status := models.WorkspaceCardStatus{}

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_card_status.Name = req.Name
		workspace_card_status.Color = req.Color

		if err := trx.Save(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "update-workspace-card-status",
			Description: fmt.Sprintf("User %s has changes to status %s", user.Username, workspace_card_status.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_log := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace card status",
		Data: structs.WorkspaceCardStatus{
			Id:    workspace_card_status.Id,
			Order: workspace_card_status.Order,
			Name:  workspace_card_status.Name,
			Color: workspace_card_status.Color,
		},
	})
}
