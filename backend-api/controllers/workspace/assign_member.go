package controllers

import (
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
	var workspace_member = models.WorkspaceMember{}
	var members = []structs.Member{}

	database.DB.First(&workspace)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	for key := range req.UserId {
		var user = models.User{}

		if err := database.DB.First(&user, req.UserId[key]).Error; err != nil {
			c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
				Success: false,
				Message: "User Not Found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		var assigned_user = models.WorkspaceMember{
			WorkspaceId: uint(req.WorkspaceId),
			UserId:      uint(req.UserId[key]),
		}

		if err := database.DB.First(&workspace_member, models.WorkspaceMember{WorkspaceId: uint(req.WorkspaceId)}).Find(&workspace_member, models.WorkspaceMember{UserId: uint(req.UserId[key])}).Error; err != nil {
			if err := database.DB.Create(&assigned_user).Error; err != nil {
				c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
					Success: false,
					Message: "Failed to assign user",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}
		} else {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "The User is already a member",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		members = append(members, structs.Member{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User assigned successfully",
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
