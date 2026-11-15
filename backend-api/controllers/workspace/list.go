package workspace

import (
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *WorkspaceController) ListWorkspace(ctx *gin.Context) {
	user_id := ctx.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := c.DB.First(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed get user info", err)
		return
	}

	response := []structs.WorkspaceResponse{}
	workspace_by_member := []uint{}

	c.DB.Model(&models.WorkspaceMember{UserId: user.Id}).Pluck("workspace_id", &workspace_by_member)

	for _, workspace_id := range workspace_by_member {
		workspace := structs.WorkspaceResponse{}
		workspace_member := []models.WorkspaceMember{}

		c.DB.Model(&models.Workspace{}).Where("id = ?", workspace_id).First(&workspace)
		c.DB.Where("workspace_id = ?", workspace_id).Preload("User").Find(&workspace_member)

		for _, member := range workspace_member {
			workspace.Members = append(workspace.Members, structs.MemberDetail{
				Id:       member.UserId,
				Name:     member.User.Name,
				Username: member.User.Username,
				Email:    member.User.Email,
				Role:     member.Role,
			})
		}

		response = append(response, structs.WorkspaceResponse{
			Id:          workspace_id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Members:     workspace.Members,
			CreatedAt:   workspace.CreatedAt,
			UpdatedAt:   workspace.UpdatedAt,
		})
	}

	helpers.SendSuccess(ctx, http.StatusOK, "Workspace retrieved successfully", response)
}
