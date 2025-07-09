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
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive workspace card status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
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

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive workspace card status",
		Data:    list_workspace_card_status,
	})
}

func ListArchivedWorkspaceCardStatus(c *gin.Context) {
	var workspace_status = []models.WorkspaceCardStatus{}
	workspace_id := c.Param("id")

	if err := database.DB.Where("is_archived = ? AND workspace_id = ?", true, workspace_id).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive archived workspace card status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
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

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive archived workspace card status",
		Data:    list_workspace_card_status,
	})
}
