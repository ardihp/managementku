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
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user_history_logs := []models.UserHistoryLog{}

	if err := database.DB.Where("user_id = ?", user_id).Order("created_at DESC").Find(&user_history_logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive history logs",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	list_histor_logs := []structs.HistoryLogResponse{}

	for _, log := range user_history_logs {
		history_log := models.HistoryLog{}

		if req.Type == "workspace" {
			database.DB.Preload(clause.Associations).Where("type = ?", req.Type).First(&history_log, log.HistoryId)
		} else {
			database.DB.Preload(clause.Associations).First(&history_log, log.HistoryId)
		}

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

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive history logs",
		Data:    list_histor_logs,
	})
}
