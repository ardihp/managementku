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
	"gorm.io/gorm/clause"
)

func UpdateWorkspaceCard(c *gin.Context) {
	req := structs.WorkspaceUpdateCardRequest{}
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	workspace_card := models.WorkspaceCard{}

	if err := database.DB.Preload(clause.Associations).First(&workspace_card, card_id).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace card not found", err)
		return
	}

	if err := database.DB.First(&models.WorkspaceCardStatus{}, req.StatusId).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace card status not found", err)
		return
	}

	workspace_card.Title = req.Title
	workspace_card.Description = req.Description
	workspace_card.WorkspaceCardStatusId = uint(req.StatusId)

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			Action:      "update-workspace-card",
			Description: fmt.Sprintf("User %s has update changes on card %s", user.Username, workspace_card.Title),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to update workspace card", err)
		return
	}

	var updated_status models.WorkspaceCardStatus
	database.DB.First(&updated_status, workspace_card.WorkspaceCardStatusId)

	helpers.SendSuccess(c, http.StatusOK, "Workspace updated successfully", structs.WorkspaceCardResponse{
		Id:          workspace_card.Id,
		Title:       workspace_card.Title,
		Description: workspace_card.Description,
		Status:      updated_status.Name,
		UserCreated: workspace_card.UserCreated.Username,
		IsArchived:  workspace_card.IsArchived,
		CreatedAt:   workspace_card.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   workspace_card.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
