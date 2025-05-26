package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspaceRole(c *gin.Context) {
	var workspace_role = []models.WorkspaceRole{}

	if err := database.DB.Find(&workspace_role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed retrieve workspace role",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var list_workspace_role = []structs.WorkspaceRoleResponse{}

	for _, workspace := range workspace_role {
		list_workspace_role = append(list_workspace_role, structs.WorkspaceRoleResponse{
			Id:   workspace.Id,
			Name: workspace.Name,

			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace role retrieved successfully",
		Data:    list_workspace_role,
	})
}
