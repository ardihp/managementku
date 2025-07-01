package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArchieveWorkspaceStatus(c *gin.Context) {
	workspace_status := models.WorkspaceStatus{}
	var id = c.Param("id")

	if err := database.DB.First(&workspace_status, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if workspace_status.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status is already archived",
			Errors:  nil,
		})
		return
	}

	workspace_status.IsArchived = true

	if err := database.DB.Save(&workspace_status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to archive workspace status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace status archived successfully",
		Data:    "",
	})
}

func UnArchieveWorkspaceStatus(c *gin.Context) {
	workspace_status := models.WorkspaceStatus{}
	var id = c.Param("id")

	if err := database.DB.First(&workspace_status, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if !workspace_status.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status is already active",
			Errors:  nil,
		})
		return
	}

	workspace_status.IsArchived = false

	if err := database.DB.Save(&workspace_status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to unarchive workspace status",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace status activated successfully",
		Data:    "",
	})
}
