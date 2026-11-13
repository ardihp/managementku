package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdateWorkspaceCardStatusOrder(c *gin.Context) {
	var req = structs.WorkspaceCardUpdateOrderStatusRequest{}
	var workspace_id = c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	for index, id := range req.OrderStatus {
		var new_workspace_card_status = models.WorkspaceCardStatus{}

		if err := database.DB.Where("id = ? AND workspace_id = ?", id, workspace_id).First(&new_workspace_card_status).Error; err != nil {
			helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace card status not found", err)
			return
		}

		new_workspace_card_status.Order = index

		if err := database.DB.Save(&new_workspace_card_status).Error; err != nil {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to update workspace order", err)
			return
		}
	}

	helpers.SendSuccess(c, http.StatusOK, "Success update workspace card status order", "")
}

func UpdateWorkspaceCardOrder(c *gin.Context) {
	req := structs.WorkspaceCardUpdateOrderRequest{}
	workspace_id := c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	if err := database.DB.Where("id = ?", workspace_id).First(&models.Workspace{}).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace not found", err)
		return
	}

	workspace_card := models.WorkspaceCard{}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		status_before := &models.WorkspaceCardStatus{}
		status_after := &models.WorkspaceCardStatus{}

		for _, order := range req.OrderStatus {
			if slices.Contains(order.CardId, req.CardId) {
				database.DB.First(&status_after, order.StatusId)
			} else {
				database.DB.First(&status_before, order.StatusId)
			}

			for index, card_id := range order.CardId {
				new_workspace_card := models.WorkspaceCard{Id: uint(card_id)}

				if err := database.DB.Where("id = ? AND workspace_id = ?", card_id, workspace_id).First(&new_workspace_card).Error; err != nil {
					return err
				}

				new_workspace_card.WorkspaceCardStatusId = uint(order.StatusId)
				new_workspace_card.Order = index

				if err := database.DB.Save(&new_workspace_card).Error; err != nil {
					return err
				}
			}
		}

		database.DB.Preload(clause.Associations).First(&workspace_card, req.CardId)

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    "manager",
			Action:      "update-order",
			Description: fmt.Sprintf("User %s has changed status for the %s card from %s to %s", user.Username, workspace_card.Title, status_before.Name, status_after.Name),
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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to update workspace card order", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Success update workspace card order", structs.WorkspaceCardResponse{
		Id:          workspace_card.Id,
		Title:       workspace_card.Title,
		Description: workspace_card.Description,
		Status:      workspace_card.WorkspaceCardStatus.Name,
		UserCreated: workspace_card.UserCreated.Username,
		IsArchived:  workspace_card.IsArchived,
		CreatedAt:   workspace_card.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   workspace_card.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
