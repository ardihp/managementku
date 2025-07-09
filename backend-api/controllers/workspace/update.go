package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdateWorkspace(c *gin.Context) {
	req := structs.WorkspaceUpdateRequest{}
	id := c.Param("id")
	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace := models.Workspace{}
	workspace_tags := []string{}

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Model(&models.Workspace{}).Preload(clause.Associations).First(&workspace, id).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace.Name = req.Name
		workspace.Description = req.Description

		if err := trx.First(&models.WorkspaceType{}, req.TypeId).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace type not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}
		workspace.Type.Id = uint(req.TypeId)

		if err := trx.First(&models.WorkspaceStatus{}, req.StatusId).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}
		workspace.Status.Id = uint(req.StatusId)

		trx.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", id).Pluck("tag", &workspace_tags)

		if len(workspace_tags) >= 1 {
			workspace_tags = []string{}
			trx.Where("workspace_id = ?", id).Delete(&models.WorkspaceTag{})
		}

		for _, tag := range req.Tags {
			var assigned_tag = models.WorkspaceTag{WorkspaceId: workspace.Id, Tag: tag}
			if err := trx.Create(&assigned_tag).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to save workspace tag",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}

			workspace_tags = append(workspace_tags, assigned_tag.Tag)
		}

		if err := trx.Save(&workspace).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "general",
			UserRole:    "manager",
			Action:      "create-workspace",
			Description: fmt.Sprintf("User %s has update changes on workspace %s", user.Username, workspace.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		user_history := models.UserHistoryLog{
			HistoryId: history_log.ID,
			UserId:    uint(user_id),
		}

		if err := trx.Create(&user_history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace updated successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Id:    workspace.Type.Id,
				Name:  workspace.Type.Name,
				Color: workspace.Type.Color,
			},
			Status: structs.WorkspaceStatus{
				Id:    workspace.Status.Id,
				Name:  workspace.Status.Name,
				Color: workspace.Status.Color,
			},
			Members:    []structs.Member{},
			Tags:       workspace_tags,
			IsArchived: workspace.IsArchived,
			CreatedAt:  workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
