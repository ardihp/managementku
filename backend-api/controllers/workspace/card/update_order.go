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
)

func UpdateWorkspaceCardStatusOrder(c *gin.Context) {
	var req = structs.WorkspaceCardUpdateOrderStatusRequest{}
	var workspace_id = c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	for index, id := range req.OrderStatus {
		var new_workspace_card_status = models.WorkspaceCardStatus{}

		if err := database.DB.Where("id = ? AND workspace_id = ?", id, workspace_id).First(&new_workspace_card_status).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace card status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		new_workspace_card_status.Order = index

		if err := database.DB.Save(&new_workspace_card_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace order",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace card status order",
		Data:    "",
	})
}

func UpdateWorkspaceCardOrder(c *gin.Context) {
	req := structs.WorkspaceCardUpdateOrderRequest{}
	workspace_id := c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("id = ?", workspace_id).First(&models.Workspace{}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		status_before := &models.WorkspaceCardStatus{}
		status_after := &models.WorkspaceCardStatus{}

		for _, order := range req.OrderStatus {
			if slices.Contains(order.CardId, req.CardId) {
				database.DB.First(&status_after, order.StatusId)
			} else {
				database.DB.First(&status_before, order.StatusId)
			}

			for index, card_id := range order.CardId {
				var new_workspace_card = models.WorkspaceCard{Id: uint(card_id)}

				if err := database.DB.Where("id = ? AND workspace_id = ?", card_id, workspace_id).First(&new_workspace_card).Error; err != nil {
					c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
						Success: false,
						Message: "Workspace card not found",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}

				new_workspace_card.Status.Id = uint(order.StatusId)
				new_workspace_card.Order = index

				if err := database.DB.Save(&new_workspace_card).Error; err != nil {
					c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
						Success: false,
						Message: "Failed to update workspace card order",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}
			}
		}

		workspace_card := models.WorkspaceCard{}
		database.DB.First(&workspace_card, req.CardId)

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    "manager",
			Action:      "update-order",
			Description: fmt.Sprintf("User %s has changed status for the %s card from %s to %s", user.Username, workspace_card.Title, status_before.Name, status_after.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace card order",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace card order",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Success update workspace card order",
			Data: structs.WorkspaceCardResponse{
				Id:          workspace_card.Id,
				Title:       workspace_card.Title,
				Description: workspace_card.Description,
				Status: structs.WorkspaceCardStatus{
					Id:         workspace_card.Status.Id,
					Order:      workspace_card.Status.Order,
					Name:       workspace_card.Status.Name,
					Color:      workspace_card.Status.Color,
					IsArchived: workspace_card.Status.IsArchived,
				},
				UserCreated: workspace_card.UserCreated.Username,
				IsArchived:  workspace_card.IsArchived,
				CreatedAt:   workspace_card.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   workspace_card.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		})

		return nil
	})
}
