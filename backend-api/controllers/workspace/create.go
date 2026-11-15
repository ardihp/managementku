package workspace

import (
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *WorkspaceController) CreateWorkspace(ctx *gin.Context) {
	req := structs.WorkspaceCreateRequest{}
	user_id := ctx.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(ctx, err)
		return
	}

	workspace := models.Workspace{
		Name:          req.Name,
		Description:   req.Description,
		UserCreatedId: user.Id,
	}

	err := c.DB.Transaction(func(trx *gorm.DB) error {
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

		workspace_member := models.WorkspaceMember{WorkspaceId: workspace.Id, UserId: uint(user_id), Role: "admin"}

		if err := trx.Create(&workspace_member).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to create workspace", err)
		return
	}

	helpers.SendSuccess(ctx, http.StatusCreated, "Workspace created successfully", structs.WorkspaceResponse{
		Id:          workspace.Id,
		Name:        workspace.Name,
		Description: workspace.Description,
		Members:     []structs.MemberDetail{},
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
	})
}
