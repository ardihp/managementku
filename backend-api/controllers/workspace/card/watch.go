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
	"gorm.io/gorm/clause"
)

func WatchWorkspaceCard(c *gin.Context) {
	req := structs.WorkspaceWatchCardRequest{}
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).Preload(clause.Associations).First(&workspace_card).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace card not found", err)
		return
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		member := models.WorkspaceCardMember{
			WorkspaceCardId: workspace_card.Id,
			UserId:          uint(req.UserId),
		}

		if err := trx.Create(&member).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			Action:      "archive-workspace-card",
			Description: fmt.Sprintf("User %s start to watch card %s", user.Username, workspace_card.Title),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to watch workspace card", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace card watched successfully", "")
}

func UnwatchWorkspaceCard(c *gin.Context) {
	req := structs.WorkspaceWatchCardRequest{}
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).Preload(clause.Associations).First(&workspace_card).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace card not found", err)
		return
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Delete(&models.WorkspaceCardMember{}, "user_id = ?", req.UserId).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			Action:      "unarchive-workspace-card",
			Description: fmt.Sprintf("User %s has unwatch card %s", user.Username, workspace_card.Title),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to unwatch workspace card", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace card unwatched successfully", "")
}
