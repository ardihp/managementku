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

func ListWorkspaceMember(c *gin.Context) {
	workspace_id := c.Param("id")

	if err := database.DB.Preload(clause.Associations).First(&models.Workspace{}, workspace_id).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace not found", err)
		return
	}

	workspace_members := []models.WorkspaceMember{}

	if err := database.DB.Where("workspace_id = ? AND role = ?", workspace_id, "member").Preload(clause.Associations).Find(&workspace_members).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Failed to retreive workspace member", err)
		return
	}

	list_workspace_members := []structs.MemberDetail{}

	for _, member := range workspace_members {
		list_workspace_members = append(list_workspace_members, structs.MemberDetail{
			Id:       member.User.Id,
			Name:     member.User.Name,
			Username: member.User.Username,
			Email:    member.User.Email,
			Role:     member.Role,
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive workspace member", list_workspace_members)
}
