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

	if err := database.DB.First(&models.WorkspaceStatus{}, req.WorkspaceStatusId).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace status not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_tags = []string{}
	database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", id).Pluck("tag", &workspace_tags)

	if len(workspace_tags) >= 1 {
		workspace_tags = []string{}
		database.DB.Where("workspace_id = ?", id).Delete(&models.WorkspaceTag{})
	}

	for _, tag := range req.Tags {
		var assigned_tag = models.WorkspaceTag{WorkspaceId: workspace.Id, Tag: tag}
		if err := database.DB.Create(&assigned_tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to save workspace tag",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		workspace_tags = append(workspace_tags, assigned_tag.Tag)
	}

	workspace.Name = req.Name
	workspace.Description = req.Description
	workspace.WorkspaceType.Id = uint(req.WorkspaceTypeId)
	workspace.WorkspaceStatus.Id = uint(req.WorkspaceStatusId)

	if err := database.DB.Save(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Model(&models.Workspace{}).Preload("WorkspaceType").First(&workspace)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace updated successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Id:    workspace.WorkspaceType.Id,
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
