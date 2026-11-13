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

func DeleteWorkspaceCardStatus(c *gin.Context) {
	workspace_card_status := models.WorkspaceCardStatus{}
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	status_id := c.Param("status_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace card status not found", err)
		return
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		workspace_cards := []models.WorkspaceCard{}
		trx.Where("workspace_id = ? AND id = ?", workspace_id, status_id).Find(&workspace_cards)

		for _, card := range workspace_cards {
			if err := trx.Where("card_id = ?", card.Id).Delete(&models.WorkspaceCardMember{}).Error; err != nil {
				return err
			}
		}

		if err := trx.Where("workspace_id = ? AND status_id = ?", workspace_id, status_id).Delete(&models.WorkspaceCard{}).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			Action:      "delete-workspace-card-status",
			Description: fmt.Sprintf("User %s has deleted status %s", user.Username, workspace_card_status.Name),
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

		if err := trx.Delete(&models.WorkspaceCardStatus{}, status_id).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to delete workspace card status", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Success delete workspace card status", nil)
}
