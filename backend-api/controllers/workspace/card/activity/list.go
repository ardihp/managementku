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

func ListWorkspaceCardActivity(c *gin.Context) {
	card_id := c.Param("card_id")

	workspace_card_activities := []models.WorkspaceCardActivity{}

	if err := database.DB.Preload(clause.Associations).Order("created_at DESC").Find(&workspace_card_activities, "card_id = ?", card_id).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve workspace card activities", err)
		return
	}

	list_workspace_card_activies := []structs.WorkspaceCardActivityResponse{}

	for _, card_activity := range workspace_card_activities {
		list_workspace_card_activies = append(list_workspace_card_activies, structs.WorkspaceCardActivityResponse{
			Id:          int(card_activity.Id),
			UserCreated: card_activity.User.Username,
			Type:        card_activity.Type,
			Message:     card_activity.Message,
			CreatedAt:   card_activity.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   card_activity.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success retrieve workspace card activities", list_workspace_card_activies)
}
