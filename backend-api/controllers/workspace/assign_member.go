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

func AssignMemberToWorkspace(c *gin.Context) {
	var req = structs.WorkspaceAssignMember{}
	var workspace = models.Workspace{}
	var role = models.WorkspaceRole{}
	var members = []structs.Member{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Model(&models.Workspace{}).Preload("WorkspaceType").First(&workspace, models.Workspace{Id: uint(req.WorkspaceId)}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	fmt.Println(workspace)

	if err := database.DB.First(&role, models.WorkspaceRole{Id: uint(req.RoleId)}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Role not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var assigned_user = []models.WorkspaceMember{}

	for key := range req.UserId {
		var user = models.User{}

		if err := database.DB.First(&user, req.UserId[key]).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "User Not Found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		if err := database.DB.Where("workspace_id = ? AND user_id = ?", req.WorkspaceId, req.UserId[key]).First(&models.WorkspaceMember{}).Error; err != nil {
			assigned_user = append(assigned_user, models.WorkspaceMember{
				WorkspaceId: uint(req.WorkspaceId),
				UserId:      uint(req.UserId[key]),
				RoleId:      uint(req.RoleId),
			})
		} else {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: fmt.Sprintf("User %s already in a member of workspace", user.Username),
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
	}

	if err := database.DB.Create(&assigned_user).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Failed to assign user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var id_member = []int{}
	database.DB.Model(&models.WorkspaceMember{}).Joins("Role").Where("workspace_id = ?", req.WorkspaceId).Order("user_id").Pluck("user_id", &id_member)

	for _, user_id := range id_member {
		var user = models.User{}
		var workspace_member = models.WorkspaceMember{}

		database.DB.Model(&models.WorkspaceMember{}).Preload("Role").Where("workspace_id = ? AND user_id = ?", req.WorkspaceId, user_id).Find(&workspace_member)

		if err := database.DB.First(&user, user_id).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "User Not Found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		members = append(members, structs.Member{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role:     workspace_member.Role.Name,
		})
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User assigned successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type:        workspace.WorkspaceType.Name,
			Members:     members,
			CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
