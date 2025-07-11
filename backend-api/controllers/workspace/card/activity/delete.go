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

func DeleteWorkspaceCardActivity(c *gin.Context) {
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	activity_id := c.Param("activity_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Where("card_id = ? AND id = ?", card_id, activity_id).Delete(&models.WorkspaceCardActivity{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card activity",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_card := models.WorkspaceCard{}
		database.DB.Find(&workspace_card, card_id)

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "delete-workspace-card-activity",
			Description: fmt.Sprintf("User %s has delete comment on card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card activity",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_log := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace card activity",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Success delete workspace card activity",
		})

		return nil
	})

}
