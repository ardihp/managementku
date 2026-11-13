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
	id := c.Param("id")
	workspace_card_statuses := []models.WorkspaceCardStatus{}

	if err := database.DB.Where("workspace_id = ?", id).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_card_statuses).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace not found", err)
		return
	}

	list_workspace_cards := []structs.WorkspaceCardByStatusResponse{}

	for _, status := range workspace_card_statuses {
		workspace_card_by_status := []models.WorkspaceCard{}

		if err := database.DB.Where("workspace_id = ? AND status_id = ? AND is_archived = ?", id, status.Id, false).Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_card_by_status).Error; err != nil {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve workspace status", err)
			return
		}

		list_workspace_card := []structs.WorkspaceCardResponse{}

		for _, card := range workspace_card_by_status {
			members := []models.WorkspaceCardMember{}

			if err := database.DB.Where("card_id = ?", card.Id).Preload(clause.Associations).Find(&members).Error; err != nil {
				helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve workspace status", err)
				return
			}

			list_member := []structs.WorkspaceMember{}

			for _, member := range members {
				list_member = append(list_member, structs.WorkspaceMember{
					Id:       member.User.Id,
					Name:     member.User.Name,
					Username: member.User.Username,
					Email:    member.User.Email,
				})
			}

			result := database.DB.Where("card_id = ? AND type = ?", card.Id, "comment").Preload(clause.Associations).Find(&[]models.WorkspaceCardActivity{})

			list_workspace_card = append(list_workspace_card, structs.WorkspaceCardResponse{
				Id:           card.Id,
				Title:        card.Title,
				Description:  card.Description,
				Status:       status.Name,
				UserCreated:  card.UserCreated.Username,
				Members:      list_member,
				CommentCount: int(result.RowsAffected),
				IsArchived:   card.IsArchived,
				IsCompleted:  card.IsCompleted,
				CreatedAt:    card.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:    card.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		list_workspace_cards = append(list_workspace_cards, structs.WorkspaceCardByStatusResponse{
			Id:         int(status.Id),
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
			Cards:      list_workspace_card,
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success retrieve workspace status", list_workspace_cards)
}

func ListArchivedWorkspaceCard(c *gin.Context) {
	workspace_id := c.Param("id")

	archived_workspace_cards := []models.WorkspaceCard{}
	list_archived_workspace_cards := []structs.WorkspaceCardResponse{}

	if err := database.DB.Where("workspace_id = ? AND is_archived = ?", workspace_id, true).Preload(clause.Associations).Find(&archived_workspace_cards).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace cards not found", err)
		return
	}

	for _, card := range archived_workspace_cards {
		members := []models.WorkspaceCardMember{}

		if err := database.DB.Where("card_id = ?", card.Id).Preload(clause.Associations).Find(&members).Error; err != nil {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to retrieve workspace card", err)
			return
		}

		list_member := []structs.WorkspaceMember{}

		for _, member := range members {
			list_member = append(list_member, structs.WorkspaceMember{
				Id:       member.User.Id,
				Name:     member.User.Name,
				Username: member.User.Username,
				Email:    member.User.Email,
			})
		}

		result := database.DB.Where("card_id = ? AND type = ?", card.Id, "comment").Preload(clause.Associations).Find(&[]models.WorkspaceCardActivity{})

		list_archived_workspace_cards = append(list_archived_workspace_cards, structs.WorkspaceCardResponse{
			Id:           card.Id,
			Title:        card.Title,
			Description:  card.Description,
			Status:       card.WorkspaceCardStatus.Name,
			UserCreated:  card.UserCreated.Username,
			Members:      list_member,
			CommentCount: int(result.RowsAffected),
			IsArchived:   card.IsArchived,
			IsCompleted:  card.IsCompleted,
			CreatedAt:    card.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    card.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive workspace card", list_archived_workspace_cards)
}
