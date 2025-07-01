package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWorkspaceStatus(c *gin.Context) {
	req := structs.WorkspaceUpdateStatusRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_status := models.WorkspaceStatus{}

	if err := database.DB.Where("id = ?", req.Id).First(&workspace_status).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_status.Name = req.Name
	workspace_status.Color = req.Color

	if err := database.DB.Save(&workspace_status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update workspace status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace status",
		Data: structs.WorkspaceStatus{
			Id:    workspace_status.Id,
			Order: workspace_status.Order,
			Name:  workspace_status.Name,
			Color: workspace_status.Color,
		},
	})
}
