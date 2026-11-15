package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspaceCard(c *gin.Context) {
	id := c.Param("id")
	workspace_card_statuses := []models.WorkspaceCardStatus{}

	if err := database.DB.Preload("Workspace").Where("workspace_id = ?", id).Order(`"order" ASC`).Find(&workspace_card_statuses).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace not found", err)
		return
	}

	response := []structs.WorkspaceCardByStatusResponse{}

	for _, status := range workspace_card_statuses {
		workspace_cards := []models.WorkspaceCard{}

		if err := database.DB.Where("workspace_id = ? AND workspace_card_status_id = ? AND is_archived = ?", id, status.Id, false).Order(`"order" ASC`).Find(&workspace_cards).Error; err != nil {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve workspace cards", err)
			return
		}

		cards := []structs.WorkspaceCardResponse{}

		for _, card := range workspace_cards {
			workspace_members := []models.WorkspaceCardMember{}

			if err := database.DB.Preload("User").Where("card_id = ?", card.Id).Find(&workspace_members).Error; err != nil {
				helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve workspace members", err)
				return
			}

			members := []structs.MemberDetail{}

			for _, member := range workspace_members {
				members = append(members, structs.MemberDetail{
					Id:       member.User.Id,
					Name:     member.User.Name,
					Username: member.User.Username,
					Email:    member.User.Email,
				})
			}

			result := database.DB.Where("card_id = ? AND type = ?", card.Id, "comment").Find(&[]models.WorkspaceCardActivity{})

			cards = append(cards, structs.WorkspaceCardResponse{
				Id:           card.Id,
				Title:        card.Title,
				Description:  card.Description,
				Status:       status.Name,
				UserCreated:  card.UserCreated.Username,
				Members:      members,
				CommentCount: int(result.RowsAffected),
				IsArchived:   card.IsArchived,
				IsCompleted:  card.IsCompleted,
				CreatedAt:    card.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    card.UpdatedAt.Format("2006-01-02 15:04:05"),
			})

			response = append(response, structs.WorkspaceCardByStatusResponse{
				Id:         int(status.Id),
				Name:       status.Name,
				Color:      status.Color,
				IsArchived: status.IsArchived,
				Cards:      cards,
			})
		}
	}

	helpers.SendSuccess(c, http.StatusOK, "Success retrieve workspace card by status", response)
}
