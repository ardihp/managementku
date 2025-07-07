package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateWorkspaceOrder(c *gin.Context) {
	req := structs.WorkspaceUpdateOrderRequest{}
	id := c.GetInt("user_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(tx *gorm.DB) error {
		status_before := &models.WorkspaceStatus{}
		status_after := &models.WorkspaceStatus{}

		for _, order := range req.OrderStatus {
			if slices.Contains(order.WorkspaceId, req.WorkspaceId) {
				database.DB.First(&status_after, order.StatusId)
			} else {
				database.DB.First(&status_before, order.StatusId)
			}

			for index, workspace := range order.WorkspaceId {

				new_workspace := models.Workspace{Id: uint(workspace)}
				tx.First(&new_workspace)

				new_workspace.Status.Id = uint(order.StatusId)
				new_workspace.Order = index

				if err := tx.Save(&new_workspace).Error; err != nil {
					c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
						Success: false,
						Message: "Failed to update workspace order",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}
			}
		}

		user := models.User{}
		tx.First(&user, id)

		workspace := models.Workspace{}
		tx.First(&workspace, req.WorkspaceId)

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "general",
			Action:      "update-order",
			Description: fmt.Sprintf("User %s has changed status for the %s workspace from %s to %s", user.Username, workspace.Name, status_before.Name, status_after.Name),
		}

		if err := tx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update workspace order",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_members := []models.WorkspaceMember{}
		tx.Where("workspace_id = ?", req.WorkspaceId).Find(&workspace_members)

		for _, member := range workspace_members {
			user_history := models.UserHistoryLog{
				HistoryId: history_log.ID,
				UserId:    member.UserId,
			}

			if err := tx.Create(&user_history).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to update workspace order",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace order",
		Data:    "",
	})
}
