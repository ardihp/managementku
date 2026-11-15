package workspace

import (
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *WorkspaceController) DetailWorkspace(ctx *gin.Context) {
	workspace_id := ctx.Param("id")

	workspace := models.Workspace{}
	workspace_members := []models.WorkspaceMember{}
	members := []structs.MemberDetail{}

	if err := c.DB.First(&workspace, workspace_id).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Workspace not found", err)
		return
	}

	if err := c.DB.Preload("User").Where("workspace_id = ?", workspace.Id).Find(&workspace_members).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to get workspace member", err)
		return
	}

	response := structs.WorkspaceDetailResponse{
		Id:          workspace.Id,
		Name:        workspace.Name,
		Description: workspace.Description,
		Members:     members,
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
	}

	for _, member := range workspace_members {
		response.Members = append(response.Members, structs.MemberDetail{
			Id:       member.User.Id,
			Name:     member.User.Name,
			Username: member.User.Username,
			Email:    member.User.Email,
			Role:     member.Role,
		})
	}

	helpers.SendSuccess(ctx, http.StatusOK, "Workspace detail retreived successfully", response)
}
