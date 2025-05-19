package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspace(c *gin.Context) {
	var list_workspace = []models.Workspace{}

	database.DB.Find(&list_workspace)

	var response = []structs.WorkspaceResponse{}

	for _, workspace := range list_workspace {
		var workspace_member = models.WorkspaceMember{}
		var members = []structs.Member{}
		var id_member = []int{}

		database.DB.Model(&workspace_member).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member)

		for _, id := range id_member {
			var user = models.User{}

			database.DB.First(&user, id)

			members = append(members, structs.Member{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     user.Role,
			})
		}

		response = append(response, structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Category:    workspace.Category,
			Members:     members,
			CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace retrieved successfully",
		Data:    response,
	})
}

func ListWorkspaceByID(c *gin.Context) {
	var workspace = models.Workspace{}
	var id = c.Param("id")

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_member = models.WorkspaceMember{}
	var members = []structs.Member{}
	var id_member = []int{}

	database.DB.Model(&workspace_member).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member)

	for _, id := range id_member {
		var user = models.User{}

		database.DB.First(&user, id)

		members = append(members, structs.Member{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace retrieved successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Category:    workspace.Category,
			Members:     members,
			CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
