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
	var req = structs.WorkspaceUpdateStatusRequest{}
	var workspace = models.Workspace{}
	var workspace_status = models.WorkspaceStatus{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.First(&workspace, models.Workspace{Id: uint(req.WorkspaceId)}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Where("workspace_id = ?", req.WorkspaceId).Delete(&workspace_status)
	var list_status = []structs.WorkspaceStatus{}

	for key := range req.Status {
		var assigned_status = models.WorkspaceStatus{
			WorkspaceId: uint(req.WorkspaceId),
			Name:        req.Status[key].Name,
			Color:       req.Status[key].Color,
		}

		if err := database.DB.Create(&assigned_status).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		list_status = append(list_status, structs.WorkspaceStatus{
			Name:  req.Status[key].Name,
			Color: req.Status[key].Color,
		})
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Workspace status update successfully",
		Data: structs.WorkspaceStatusResponse{
			WorkspaceId: workspace.Id,
			Status:      list_status,
		},
	})
}
