package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArchieveWorkspace(c *gin.Context) {
	var workspace = models.Workspace{}
	var id = c.Param("id")

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if workspace.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace is already archived",
			Errors:  nil,
		})
		return
	}

	workspace.IsArchived = true

	if err := database.DB.Save(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to archive workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace archived successfully",
		Data:    "",
	})
}

func UnArchieveWorkspace(c *gin.Context) {
	var workspace = models.Workspace{}
	var id = c.Param("id")

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if !workspace.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace is already active",
			Errors:  nil,
		})
		return
	}

	workspace.IsArchived = false

	if err := database.DB.Save(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to unarchive workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace activated successfully",
		Data:    "",
	})
}
