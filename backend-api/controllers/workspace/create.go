package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkspace(c *gin.Context) {
	req := structs.WorkspaceCreateRequest{}
	user_id := c.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	workspace := models.Workspace{
		Name:          req.Name,
		Description:   req.Description,
		UserCreatedId: user.Id,
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Create(&workspace).Error; err != nil {
			return err
		}

		statuses := []structs.WorkspaceCardStatus{
			{Name: "Backlog", Color: "gray", Order: 0},
			{Name: "In Progress", Color: "yellow", Order: 1},
			{Name: "Completed", Color: "green", Order: 2}}

		for _, status := range statuses {
			workspace_card_status := models.WorkspaceCardStatus{
				WorkspaceId: workspace.Id,
				Name:        status.Name,
				Color:       status.Color,
				Order:       status.Order,
			}

			if err := trx.Create(&workspace_card_status).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to create workspace", err)
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Workspace created successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Members:     []structs.WorkspaceMember{},
			CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
