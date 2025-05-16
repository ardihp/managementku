package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWorkspace(c *gin.Context) {
	var req = structs.WorkspaceCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace := models.Workspace{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
	}

	if err := database.DB.Create(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_member = []models.WorkspaceMember{}
	var members = []models.User{}

	if len(req.Members) >= 1 {
		for key := range req.Members {
			var user = models.User{}

			if err := database.DB.First(&user, req.Members[key]).Error; err != nil {
				c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
					Success: false,
					Message: "User Not Found",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}

			workspace_member = append(workspace_member, models.WorkspaceMember{WorkspaceId: int(workspace.Id), UserId: int(req.Members[key])})
			members = append(members, user)
		}

		if err := database.DB.Create(&workspace_member).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to assign member to workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Workspace created successfully",
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
