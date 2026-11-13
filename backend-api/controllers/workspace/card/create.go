package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkspaceCard(c *gin.Context) {
	var req = structs.WorkspaceCreateCardRequest{}
	var workspace_id = c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	if err := database.DB.First(&models.Workspace{}, workspace_id).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Workspace not found", err)
		return
	}

	workspace_card_status := models.WorkspaceCardStatus{}

	if err := database.DB.Where("id = ? AND workspace_id = ? AND is_archived = ?", req.StatusId, workspace_id, false).First(&workspace_card_status).Error; err != nil {
		helpers.SendError(c, http.StatusNotFound, "Status card not found", err)
		return
	}

	var card models.WorkspaceCard

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		result := trx.Where("status_id = ?", req.StatusId).Find(&[]models.WorkspaceCard{})

		card = models.WorkspaceCard{
			WorkspaceId:           uint(parsed_workspace_id),
			WorkspaceCardStatusId: uint(req.StatusId),
			UserCreatedId:         uint(user_id),
			Title:                 req.Title,
			Description:           req.Description,
			Order:                 int(result.RowsAffected),
		}

		if err := trx.Create(&card).Error; err != nil {
			return err
		}

		workspace_card_activity := models.WorkspaceCardActivity{
			WorkspaceCardId: card.Id,
			UserId:          uint(user_id),
			Type:            "general",
			Message:         fmt.Sprintf("User %s created this card", user.Username),
		}

		if err := trx.Create(&workspace_card_activity).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			Action:      "create-workspace-card",
			Description: fmt.Sprintf("User %s has added new card to status %s", user.Username, workspace_card_status.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			return err
		}

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to create workspace card", err)
		return
	}

	helpers.SendSuccess(c, http.StatusCreated, "Workspace card created successfully", structs.WorkspaceCardResponse{
		Id:          card.Id,
		Title:       card.Title,
		Description: card.Description,
		Members:     []structs.WorkspaceMember{},
	})
}
