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
	var list_workspace = []int{}
	var user_id = c.GetInt("user_id")

	database.DB.Model(&models.WorkspaceMember{}).Where("user_id = ?", user_id).Pluck("workspace_id", &list_workspace)

	var response = []structs.WorkspaceResponse{}

	for _, workspace_id := range list_workspace {
		var members = []structs.Member{}
		var id_member = []int{}
		var workspace = models.Workspace{
			Id: uint(workspace_id),
		}

		database.DB.First(&workspace)
		database.DB.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace_id).Order("user_id").Pluck("user_id", &id_member)

		for _, user_id := range id_member {
			var user = models.User{}
			var workspace_member = models.WorkspaceMember{}

			database.DB.First(&user, user_id)
			database.DB.Model(&models.WorkspaceMember{}).Preload("Role").Where("workspace_id = ? AND user_id = ?", workspace_id, user_id).Find(&workspace_member)

			members = append(members, structs.Member{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     workspace_member.Role.Name,
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
