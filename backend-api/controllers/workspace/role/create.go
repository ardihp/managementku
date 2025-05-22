package controller

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWorkspaceRole(c *gin.Context) {
	req := structs.WorkspaceCreateRoleRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var role = models.WorkspaceRole{
		Name: req.Name,
	}

	if err := database.DB.Create(&role).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace role",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace role retreived successfully",
		Data: structs.WorkspaceRoleReponse{
			Id:        role.Id,
			Name:      role.Name,
			CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
