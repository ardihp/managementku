package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspaceStatus(c *gin.Context) {
	var id = c.Param("id")
	var is_archived = c.Query("is_archived")
	var workspace = []models.WorkspaceStatus{}

	if err := database.DB.Where("workspace_id = ? && is_archived = ?", id, is_archived == "true").Find(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_status = []structs.WorkspaceStatusCardResponse{}

	for _, status := range workspace {
		var workspace_status_card = []models.WorkspaceCard{}

		if err := database.DB.Where("workspace_id = ? && status = ?", id, status.Name).Find(&workspace_status_card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to retrieve workspace status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		var list_workspace_status_card = []structs.WorkspaceCardResponse{}

		for _, card := range workspace_status_card {
			var members = []models.WorkspaceMember{}

			if err := database.DB.Where("workspace_id = ?", id).Preload("User").Find(&members).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to retrieve workspace status",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}

			var list_member = []structs.Member{}

			for _, member := range members {
				list_member = append(list_member, structs.Member{
					Id:       member.User.Id,
					Name:     member.User.Name,
					Username: member.User.Username,
					Email:    member.User.Email,
				})
			}

			list_workspace_status_card = append(list_workspace_status_card, structs.WorkspaceCardResponse{
				Id:          card.Id,
				Title:       card.Title,
				Description: card.Description,
				Members:     list_member,
			})
		}

		workspace_status = append(workspace_status, structs.WorkspaceStatusCardResponse{
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
			Cards:      list_workspace_status_card,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success retrieve workspace status",
		Data:    workspace_status,
	})
}
