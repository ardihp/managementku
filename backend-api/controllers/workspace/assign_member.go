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
	var members = []models.User{}

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

		members = append(members, user)

		var workspace_member = models.WorkspaceMember{
			WorkspaceId: int(req.WorkspaceId),
			UserId:      int(req.UserId[key]),
		}

		if err := database.DB.Where("workspace_id ?", workspace_member.WorkspaceId).Find(&workspace_member).Error; err != nil {
			if helpers.IsDuplicateEntryError(err) {
				c.JSON(http.StatusConflict, structs.ErrorResponse{
					Success: false,
					Message: "Duplicate entry error",
					Errors:  helpers.TranslateErrorMessage(err),
				})
			} else {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to assign user",
					Errors:  helpers.TranslateErrorMessage(err),
				})
			}
			return
		}

		if err := database.DB.Create(&workspace_member).Error; err != nil {
			c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
				Success: false,
				Message: "Failed to assign user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
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
