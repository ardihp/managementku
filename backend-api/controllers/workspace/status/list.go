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

func ListWorkspaceStatus(c *gin.Context) {
	var workspace_status = []models.WorkspaceStatus{}

	if err := database.DB.Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive workspace status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var list_workspace_status = []structs.WorkspaceStatus{}

	for _, status := range workspace_status {
		list_workspace_status = append(list_workspace_status, structs.WorkspaceStatus{
			Id:         status.Id,
			Order:      status.Order,
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive workspace status",
		Data:    list_workspace_status,
	})
}
