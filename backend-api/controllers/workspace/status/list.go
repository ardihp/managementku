package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspaceStatus(c *gin.Context) {
	var id = c.Param("id")
	var workspace = []models.WorkspaceStatus{}

	if err := database.DB.Where("workspace_id = ?", id).Find(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_status = []structs.WorkspaceStatus{}

	for _, status := range workspace {
		workspace_status = append(workspace_status, structs.WorkspaceStatus{
			Name:  status.Name,
			Color: status.Color,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success retrieve workspace status",
		Data:    workspace_status,
	})
}
