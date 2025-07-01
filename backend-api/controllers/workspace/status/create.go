package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWorkspaceStatus(c *gin.Context) {
	var req = structs.WorkspaceCreateStatusRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if result := database.DB.Where("name = ?", req.Name).Find(&models.WorkspaceStatus{}); result.RowsAffected >= 1 {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status already exists",
		})
		return
	}

	list_workspace_status := []models.WorkspaceStatus{}
	database.DB.Find(&list_workspace_status)

	var workspace_status = models.WorkspaceStatus{
		Order: len(list_workspace_status) + 1,
		Name:  req.Name,
		Color: req.Color,
	}

	if err := database.DB.Create(&workspace_status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create workspace status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive workspace status",
		Data: structs.WorkspaceStatus{
			Id:         workspace_status.Id,
			Order:      workspace_status.Order,
			Name:       workspace_status.Name,
			Color:      workspace_status.Color,
			IsArchived: workspace_status.IsArchived,
		},
	})
}
