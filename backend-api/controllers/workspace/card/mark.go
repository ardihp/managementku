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

func MarkWorkspaceCardComplete(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Preload(clause.Associations).First(&workspace_card, card_id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		workspace_card.IsCompleted = true

		if err := trx.Save(&workspace_card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to complete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "complete-workspace-card",
			Description: fmt.Sprintf("User %s mark complete to card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to complete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to complete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Workspace card completed successfully",
			Data:    "",
		})

		return nil
	})
}

func MarkWorkspaceCardIncomplete(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Preload(clause.Associations).First(&workspace_card, card_id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		workspace_card.IsCompleted = false

		if err := trx.Save(&workspace_card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to incomplete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "incomplete-workspace-card",
			Description: fmt.Sprintf("User %s mark incomplete to card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to incomplete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		parsed_workspace_id, _ := strconv.Atoi(workspace_id)

		workspace_history := models.WorkspaceHistoryLog{
			HistoryId:   history_log.ID,
			WorkspaceId: uint(parsed_workspace_id),
		}

		if err := trx.Create(&workspace_history).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to incomplete workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Workspace card incomplete successfully",
			Data:    "",
		})

		return nil
	})
}
