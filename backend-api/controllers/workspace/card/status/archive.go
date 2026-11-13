package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArchieveWorkspaceCardStatus(c *gin.Context) {
	workspace_card_status := models.WorkspaceCardStatus{}
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	status_id := c.Param("status_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace card status not found", err)
		return
	}

	if workspace_card_status.IsArchived {
		helpers.SendError(c, http.StatusConflict, "Workspace card status is already archived", nil)
		return
	}

	workspace_card_status.IsArchived = true

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card_status).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			Action:      "archive-workspace-card-status",
			Description: fmt.Sprintf("User %s has archived status %s", user.Username, workspace_card_status.Name),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to archive workspace card status", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace card status archived successfully", "")
}

func UnArchieveWorkspaceCardStatus(c *gin.Context) {
	workspace_card_status := models.WorkspaceCardStatus{}
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	status_id := c.Param("status_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace card status not found", err)
		return
	}

	if !workspace_card_status.IsArchived {
		helpers.SendError(c, http.StatusConflict, "Workspace card status is already active", nil)
		return
	}

	workspace_card_status.IsArchived = false

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card_status).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			Action:      "unarchive-workspace-card-status",
			Description: fmt.Sprintf("User %s has unarchive status %s", user.Username, workspace_card_status.Name),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to unarchive workspace card status", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace card status activated successfully", "")
}
