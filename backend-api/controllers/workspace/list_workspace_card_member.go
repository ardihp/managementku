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

func ListWorkspaceCardMember(c *gin.Context) {
	card_id := c.Param("id")

	if err := database.DB.Preload(clause.Associations).First(&models.WorkspaceCard{}, card_id).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace card not found", err)
		return
	}

	workspace_card_members := []models.WorkspaceCardMember{}

	if err := database.DB.Where("workspace_card_id = ?", card_id).Preload(clause.Associations).Find(&workspace_card_members).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Failed to retreive workspace card member", err)
		return
	}

	list_workspace_card_members := []structs.MemberDetail{}

	for _, member := range workspace_card_members {
		list_workspace_card_members = append(list_workspace_card_members, structs.MemberDetail{
			Id:       member.User.Id,
			Name:     member.User.Name,
			Username: member.User.Username,
			Email:    member.User.Email,
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive workspace card member", list_workspace_card_members)
}
