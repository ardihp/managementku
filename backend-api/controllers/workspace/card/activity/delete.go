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

func DeleteWorkspaceCardActivity(c *gin.Context) {
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	activity_id := c.Param("activity_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Where("card_id = ? AND id = ?", card_id, activity_id).Delete(&models.WorkspaceCardActivity{}).Error; err != nil {
			return err
		}

		workspace_card := models.WorkspaceCard{}
		database.DB.Find(&workspace_card, card_id)

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			Action:      "delete-workspace-card-activity",
			Description: fmt.Sprintf("User %s has delete comment on card %s", user.Username, workspace_card.Title),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to delete workspace card activity", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Success delete workspace card activity", nil)
}
