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
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_card_members := []models.WorkspaceCardMember{}

	if err := database.DB.Where("card_id = ?", card_id).Preload(clause.Associations).Find(&workspace_card_members).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive workspace card member",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	list_workspace_card_members := []structs.Member{}

	for _, member := range workspace_card_members {
		list_workspace_card_members = append(list_workspace_card_members, structs.Member{
			Id:       member.User.Id,
			Name:     member.User.Name,
			Username: member.User.Username,
			Email:    member.User.Email,
			Role:     member.User.Role,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive workspace card member",
		Data:    list_workspace_card_members,
	})
}
