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

func CreateWorkspace(c *gin.Context) {
	req := structs.WorkspaceCreateRequest{}
	user_id := c.GetInt("user_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace := models.Workspace{
		Name:        req.Name,
		Description: req.Description,
		TypeId:      uint(req.TypeId),
		StatusId:    uint(req.StatusId),
	}
	workspace_tags := []string{}
	user := models.User{}
	database.DB.First(&user, user_id)

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.First(&models.WorkspaceType{}, req.TypeId).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace type not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.First(&models.WorkspaceStatus{}, req.StatusId).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		result := trx.Model(&models.Workspace{}).Where("status_id = ?", req.StatusId)
		workspace.Order = int(result.RowsAffected)

		if err := trx.Create(&workspace).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		for _, tag := range req.Tags {
			assigned_tag := models.WorkspaceTag{WorkspaceId: workspace.Id, Tag: tag}

			if err := trx.Create(&assigned_tag).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to add workspace tag",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}

			workspace_tags = append(workspace_tags, assigned_tag.Tag)
		}

		if err := trx.Model(&models.Workspace{}).Preload(clause.Associations).First(&workspace, &models.Workspace{TypeId: uint(req.TypeId)}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to read workspace type",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		var assigned_user = models.WorkspaceMember{
			WorkspaceId: workspace.Id,
			UserId:      uint(user_id),
			Role:        "admin",
		}

		if err := trx.Create(&assigned_user).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Failed to assign user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		var workspace_card_status = []structs.WorkspaceCardStatus{
			{Name: "Backlog", Color: "gray", Order: 0},
			{Name: "In Progress", Color: "yellow", Order: 1},
			{Name: "Completed", Color: "green", Order: 2}}

		var default_card = []models.WorkspaceCardStatus{}

		for _, status := range workspace_card_status {
			default_card = append(default_card, models.WorkspaceCardStatus{
				WorkspaceId: workspace.Id,
				Name:        status.Name,
				Color:       status.Color,
				Order:       status.Order,
			})
		}

		if err := trx.Create(&default_card).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create workspace status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "general",
			UserRole:    "manager",
			Action:      "create-workspace",
			Description: fmt.Sprintf("User %s has created new workspace at status %s", user.Username, workspace.Status.Name),
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

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Workspace created successfully",
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
				Id:         workspace.Status.Id,
				Name:       workspace.Status.Name,
				Color:      workspace.Status.Color,
				IsArchived: workspace.Status.IsArchived,
			},
			Members:   []structs.Member{},
			Tags:      workspace_tags,
			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
