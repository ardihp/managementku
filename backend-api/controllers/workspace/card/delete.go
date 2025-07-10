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

func DeleteWorkspaceCard(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).First(&models.WorkspaceCard{}).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		workspace_card := models.WorkspaceCard{}
		trx.Where("workspace_id = ? AND id = ?", workspace_id, card_id).First(&workspace_card)

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "delete-workspace-card",
			Description: fmt.Sprintf("User %s has deleted card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card",
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
				Message: "Failed to delete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("card_id = ?", card_id).Delete(&models.WorkspaceCardActivity{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("card_id = ?", card_id).Delete(&models.WorkspaceCardMember{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("workspace_id = ? AND id = ?", workspace_id, card_id).Delete(&models.WorkspaceCard{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Workspace card deleted successfully",
			Data:    "",
		})

		return nil
	})
}
