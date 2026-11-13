package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListHistoryLog(c *gin.Context) {
	req := structs.ListHistoryLogRequest{}
	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindQuery(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	list_histor_logs := []structs.HistoryLogResponse{}
	user_history_logs := []models.UserHistoryLog{}

	if err := database.DB.Where("user_id = ?", user_id).Order("created_at DESC").Find(&user_history_logs).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retreive history logs", err)
		return
	}

	for _, log := range user_history_logs {
		history_log := models.HistoryLog{}

		database.DB.Preload(clause.Associations).First(&history_log, log.HistoryId)

		list_histor_logs = append(list_histor_logs, structs.HistoryLogResponse{
			ID:          int(history_log.ID),
			User:        history_log.User.Username,
			Type:        history_log.Type,
			Action:      history_log.Action,
			Description: history_log.Description,
			CreatedAt:   history_log.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   history_log.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive history logs", list_histor_logs)
}

func ListWorkspaceHistoryLog(c *gin.Context) {
	workspace_id := c.Param("id")

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	list_history_logs := []structs.HistoryLogResponse{}

	workspace_history_logs := []models.WorkspaceHistoryLog{}

	if err := database.DB.Where("workspace_id = ?", workspace_id).Order("created_at DESC").Find(&workspace_history_logs).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retreive history logs", err)
		return
	}

	for _, log := range workspace_history_logs {
		history_log := models.HistoryLog{}

		database.DB.Preload(clause.Associations).First(&history_log, log.HistoryId)

		list_history_logs = append(list_history_logs, structs.HistoryLogResponse{
			ID:          int(history_log.ID),
			User:        history_log.User.Username,
			Type:        history_log.Type,
			Action:      history_log.Action,
			Description: history_log.Description,
			CreatedAt:   history_log.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   history_log.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive workspace history logs", list_history_logs)
}
