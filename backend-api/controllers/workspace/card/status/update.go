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
		helpers.ValidationError(c, err)
		return
	}

	workspace_card_status := models.WorkspaceCardStatus{}

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace status not found", err)
		return
	}

	workspace_card_status.Name = req.Name
	workspace_card_status.Color = req.Color

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card_status).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			Action:      "update-workspace-card-status",
			Description: fmt.Sprintf("User %s has changes to status %s", user.Username, workspace_card_status.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_log := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_log).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to update workspace card status", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Success update workspace card status", structs.WorkspaceCardStatus{
		Id:    workspace_card_status.Id,
		Order: workspace_card_status.Order,
		Name:  workspace_card_status.Name,
		Color: workspace_card_status.Color,
	})
}
