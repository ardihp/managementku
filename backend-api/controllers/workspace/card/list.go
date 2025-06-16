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

func ListWorkspaceCard(c *gin.Context) {
	var id = c.Param("id")
	var workspace_cards = []models.WorkspaceCardStatus{}

	if err := database.DB.Where("workspace_id = ? AND is_archived = ?", id, false).Find(&workspace_cards).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var list_workspace_cards = []structs.WorkspaceCardByStatusResponse{}

	for _, status := range workspace_cards {
		var workspace_card_by_status = []models.WorkspaceCard{}

		if err := database.DB.Where("workspace_id = ? AND status_id = ?", id, status.Id).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}}).Find(&workspace_card_by_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to retrieve workspace status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		var list_workspace_card_by_status = []structs.WorkspaceCardResponse{}

		for _, card := range workspace_card_by_status {
			var members = []models.WorkspaceCardMember{}

			if err := database.DB.Where("card_id = ?", card.Id).Preload(clause.Associations).Find(&members).Error; err != nil {
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

			list_workspace_card_by_status = append(list_workspace_card_by_status, structs.WorkspaceCardResponse{
				Id:          card.Id,
				Title:       card.Title,
				Description: card.Description,
				Members:     list_member,
			})
		}

		list_workspace_cards = append(list_workspace_cards, structs.WorkspaceCardByStatusResponse{
			Id:         int(status.Id),
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
			Cards:      list_workspace_card_by_status,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success retrieve workspace status",
		Data:    list_workspace_cards,
	})
}
