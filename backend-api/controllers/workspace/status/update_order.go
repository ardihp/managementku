package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOrderWorkspaceStatus(c *gin.Context) {
	req := structs.WorkspaceUpdateTypeOrderRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	for index, id := range req.OrderStatus {
		new_workspace_card_status := models.WorkspaceStatus{}

		if err := database.DB.Where("id = ?", id).First(&new_workspace_card_status).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}

		new_workspace_card_status.Order = index

		if err := database.DB.Save(&new_workspace_card_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace status order",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace status order",
		Data:    "",
	})
}
