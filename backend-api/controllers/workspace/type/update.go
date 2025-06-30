package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWorkspaceType(c *gin.Context) {
	req := structs.WorkspaceUpdateTypeRequest{}

	workspace_type := models.WorkspaceType{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.First(&workspace_type, req.Id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace type not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_type.Name = req.Name
	workspace_type.Color = req.Color

	if err := database.DB.Where("id = ?", req.Id).Save(&workspace_type).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed update workspace type",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace type update successfully",
		Data: structs.WorkspaceTypeResponse{
			Id:        workspace_type.Id,
			Name:      workspace_type.Name,
			Color:     workspace_type.Color,
			CreatedAt: workspace_type.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace_type.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
