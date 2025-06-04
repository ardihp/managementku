package controllers

import (
	"fmt"
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
	var user = models.User{Id: uint(user_id)}

	if err := database.DB.First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed get user info",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if user.Role == "employee" {
		database.DB.Model(&models.WorkspaceMember{}).Where("user_id = ?", user_id).Pluck("workspace_id", &list_workspace)
	} else {
		database.DB.Model(&models.Workspace{}).Pluck("id", &list_workspace)
	}

	fmt.Println(list_workspace)

	var response = []structs.WorkspaceResponse{}

	for _, workspace_id := range list_workspace {
		var members = []structs.Member{}
		var id_member = []int{}
		var workspace = models.Workspace{
			Id: uint(workspace_id),
		}
		var workspace_tags = []string{}

		database.DB.Model(&models.Workspace{}).Preload("WorkspaceType").First(&workspace)
		database.DB.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace_id).Order("user_id").Pluck("user_id", &id_member)
		database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", workspace_id).Pluck("tag", &workspace_tags)

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
			Type: structs.WorkspaceType{
				Name:  workspace.WorkspaceType.Name,
				Color: workspace.WorkspaceType.Color,
			},
			Members:   members,
			Tags:      workspace_tags,
			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
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
	var workspace_id = c.Param("id")
	var workspace_tags = []string{}

	if err := database.DB.Model(&models.Workspace{}).Preload("WorkspaceType").First(&workspace, workspace_id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var members = []structs.Member{}
	var id_member = []int{}

	database.DB.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member)
	database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", workspace_id).Pluck("tag", &workspace_tags)

	for _, user_id := range id_member {
		var user = models.User{}
		var workspace_member = models.WorkspaceMember{}

		database.DB.First(&user, user_id)
		database.DB.Model(&models.WorkspaceMember{}).Preload("Role").Where("workspace_id = ? AND user_id = ?", workspace_id, user_id).First(&workspace_member)

		members = append(members, structs.Member{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role:     workspace_member.Role.Name,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace retrieved successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Name:  workspace.WorkspaceType.Name,
				Color: workspace.WorkspaceType.Color,
			},
			Members:   members,
			Tags:      workspace_tags,
			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
