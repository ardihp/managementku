package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWorkspaceType(c *gin.Context) {
	var req = structs.WorkspaceCreateTypeRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_type = models.WorkspaceType{
		Name:  req.Name,
		Color: req.Color,
	}

	if err := database.DB.Create(&workspace_type).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed create workspace type",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace type created successfully",
		Data: structs.WorkspaceTypeResponse{
			Id:        workspace_type.Id,
			Name:      workspace_type.Name,
			Color:     workspace_type.Color,
			CreatedAt: workspace_type.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace_type.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
