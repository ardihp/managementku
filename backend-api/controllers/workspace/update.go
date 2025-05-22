package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWorkspace(c *gin.Context) {
	var req = structs.WorkspaceUpdateRequest{}
	var workspace = models.Workspace{}
	var id = c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Model(&models.Workspace{}).Preload("WorkspaceType").First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.First(&models.WorkspaceType{}, req.WorkspaceTypeId).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace type not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace.Name = req.Name
	workspace.Description = req.Description
	workspace.WorkspaceTypeId = uint(req.WorkspaceTypeId)

	if err := database.DB.Save(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Model(&models.Workspace{}).Preload("WorkspaceType").First(&workspace, &models.Workspace{WorkspaceTypeId: uint(req.WorkspaceTypeId)})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace updated successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type:        workspace.WorkspaceType.Name,
			Members:     []structs.Member{},
			CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
