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

func ArchieveWorkspaceCard(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).First(&workspace_card).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if workspace_card.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace is already archived",
			Errors:  nil,
		})
		return
	}

	workspace_card.IsArchived = true

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_card_activity := models.WorkspaceCardActivity{
			CardId:  workspace_card.Id,
			UserId:  uint(user_id),
			Type:    "general",
			Message: fmt.Sprintf("User %s archive this card", user.Username),
		}

		if err := trx.Create(&workspace_card_activity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "archive-workspace-card",
			Description: fmt.Sprintf("User %s has archived card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace",
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
				Message: "Failed to archive workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Workspace card archived successfully",
			Data:    "",
		})

		return nil
	})
}

func UnarchieveWorkspaceCard(c *gin.Context) {
	workspace_id := c.Param("id")
	card_id := c.Param("card_id")
	workspace_card := models.WorkspaceCard{}

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.Where("workspace_id = ? AND id = ?", workspace_id, card_id).First(&workspace_card).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if !workspace_card.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace is already active",
			Errors:  nil,
		})
		return
	}

	workspace_card.IsArchived = false

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace_card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_card_activity := models.WorkspaceCardActivity{
			CardId:  workspace_card.Id,
			UserId:  uint(user_id),
			Type:    "general",
			Message: fmt.Sprintf("User %s unarchive this card", user.Username),
		}

		if err := trx.Create(&workspace_card_activity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "workspace",
			UserRole:    user.Role,
			Action:      "unarchive-workspace-card",
			Description: fmt.Sprintf("User %s has unarchived card %s", user.Username, workspace_card.Title),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace card",
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
				Message: "Failed to unarchive workspace card",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		c.JSON(http.StatusOK, structs.SuccessResponse{
			Success: true,
			Message: "Workspace card activated successfully",
			Data:    "",
		})

		return nil
	})
}
