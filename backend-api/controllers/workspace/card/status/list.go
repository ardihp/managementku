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

func ListWorkspaceCardStatus(c *gin.Context) {
	workspace_status := []models.WorkspaceCardStatus{}
	workspace_id := c.Param("id")

	if err := database.DB.Where("is_archived = ? AND workspace_id = ?", false, workspace_id).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_status).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retreive workspace card status", err)
		return
	}

	var list_workspace_card_status = []structs.WorkspaceCardStatus{}

	for _, status := range workspace_status {
		list_workspace_card_status = append(list_workspace_card_status, structs.WorkspaceCardStatus{
			Id:         status.Id,
			Order:      status.Order,
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive workspace card status", list_workspace_card_status)
}

func ListArchivedWorkspaceCardStatus(c *gin.Context) {
	var workspace_status = []models.WorkspaceCardStatus{}
	workspace_id := c.Param("id")

	if err := database.DB.Where("is_archived = ? AND workspace_id = ?", true, workspace_id).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_status).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to retreive archived workspace card status", err)
		return
	}

	var list_workspace_card_status = []structs.WorkspaceCardStatus{}

	for _, status := range workspace_status {
		list_workspace_card_status = append(list_workspace_card_status, structs.WorkspaceCardStatus{
			Id:         status.Id,
			Order:      status.Order,
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive archived workspace card status", list_workspace_card_status)
}
