package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspaceType(c *gin.Context) {
	var workspace_type = []models.WorkspaceType{}

	if err := database.DB.Find(&workspace_type).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed retrieve workspace type",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var list_workspace_type = []structs.WorkspaceTypeResponse{}

	for _, workspace := range workspace_type {
		list_workspace_type = append(list_workspace_type, structs.WorkspaceTypeResponse{
			Id:        workspace.Id,
			Name:      workspace.Name,
			Color:     workspace.Color,
			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace type retrieved successfully",
		Data:    list_workspace_type,
	})
}
