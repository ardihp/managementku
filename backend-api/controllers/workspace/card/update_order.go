package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var req = structs.WorkspaceCardUpdateOrderRequest{}
	var workspace_id = c.Param("id")

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

	for _, order := range req.OrderStatus {
		for index, card_id := range order.CardId {
			var new_workspace_card = models.WorkspaceCard{Id: uint(card_id)}

			if err := database.DB.Where("id = ? AND workspace_id = ?", card_id, workspace_id).First(&new_workspace_card).Error; err != nil {
				c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
					Success: false,
					Message: "Workspace card not found",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}

			new_workspace_card.Status.Id = uint(order.StatusId)
			new_workspace_card.Order = index

			if err := database.DB.Save(&new_workspace_card).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to update workspace card order",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace card order",
		Data:    "",
	})
}
