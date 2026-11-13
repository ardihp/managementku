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

func ArchieveWorkspaceCard(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).First(&workspace_card).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace card not found", err)
		return
	}

	if workspace_card.IsArchived {
		helpers.SendError(c, http.StatusConflict, "Workspace is already archived", nil)
		return
	}

	workspace_card.IsArchived = true

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card).Error; err != nil {
			return err
		}

		workspace_card_activity := models.WorkspaceCardActivity{
			WorkspaceCardId: workspace_card.Id,
			UserId:          uint(user_id),
			Type:            "general",
			Message:         fmt.Sprintf("User %s archive this card", user.Username),
		}

		if err := trx.Create(&workspace_card_activity).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			Action:      "archive-workspace-card",
			Description: fmt.Sprintf("User %s has archived card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to archive workspace card", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace card archived successfully", "")
}

func UnarchieveWorkspaceCard(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).First(&workspace_card).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace card not found", err)
		return
	}

	if !workspace_card.IsArchived {
		helpers.SendError(c, http.StatusConflict, "Workspace is already active", nil)
		return
	}

	workspace_card.IsArchived = false

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card).Error; err != nil {
			return err
		}

		workspace_card_activity := models.WorkspaceCardActivity{
			WorkspaceCardId: workspace_card.Id,
			UserId:          uint(user_id),
			Type:            "general",
			Message:         fmt.Sprintf("User %s unarchive this card", user.Username),
		}

		if err := trx.Create(&workspace_card_activity).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			Action:      "unarchive-workspace-card",
			Description: fmt.Sprintf("User %s has unarchived card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to unarchive workspace card", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace card activated successfully", "")
}
