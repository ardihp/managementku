package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	if err := database.DB.First(&models.WorkspaceType{}, req.WorkspaceTypeId).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace type not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace := models.Workspace{
		Name:            req.Name,
		Description:     req.Description,
		WorkspaceTypeId: uint(req.WorkspaceTypeId),
	}

	if err := database.DB.Create(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_tags := []string{}

	for _, tag := range req.Tags {
		var assigned_tag = models.WorkspaceTag{WorkspaceId: workspace.Id, Tag: tag}
		if err := database.DB.Create(&assigned_tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to add workspace tag",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		workspace_tags = append(workspace_tags, assigned_tag.Tag)
	}

	if err := database.DB.Model(&models.Workspace{}).Preload(clause.Associations).First(&workspace, &models.Workspace{WorkspaceTypeId: uint(req.WorkspaceTypeId)}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to read workspace type",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Workspace created successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Name:  workspace.WorkspaceType.Name,
				Color: workspace.WorkspaceType.Color,
			},
			Members:   []structs.Member{},
			Tags:      workspace_tags,
			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
