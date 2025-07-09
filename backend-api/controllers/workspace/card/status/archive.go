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

func ArchieveWorkspaceCardStatus(c *gin.Context) {
	workspace_card_status := models.WorkspaceCardStatus{}
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	status_id := c.Param("status_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusNotFound, structs.ErrorResponse{
				Success: false,
				Message: "Workspace card status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if workspace_card_status.IsArchived {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Workspace card status is already archived",
				Errors:  nil,
			})
			return trx.Error
		}

		workspace_card_status.IsArchived = true

		if err := trx.Save(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			UserRole:    "manager",
			Action:      "archive-workspace-card-status",
			Description: fmt.Sprintf("User %s has archived status %s", user.Username, workspace_card_status.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace card status",
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
				Message: "Failed to archive workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace card status archived successfully",
		Data:    "",
	})
}

func UnArchieveWorkspaceCardStatus(c *gin.Context) {
	workspace_card_status := models.WorkspaceCardStatus{}
	user_id := c.GetInt("user_id")
	workspace_id := c.Param("id")
	status_id := c.Param("status_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Where("workspace_id = ? AND id = ?", workspace_id, status_id).First(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusNotFound, structs.ErrorResponse{
				Success: false,
				Message: "Workspace card status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if !workspace_card_status.IsArchived {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Workspace card status is already active",
				Errors:  nil,
			})
			return trx.Error
		}

		workspace_card_status.IsArchived = false

		if err := trx.Save(&workspace_card_status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      uint(user_id),
			Type:        "workspace",
			UserRole:    "manager",
			Action:      "unarchive-workspace-card-status",
			Description: fmt.Sprintf("User %s has unarchive status %s", user.Username, workspace_card_status.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace card status",
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
				Message: "Failed to unarchive workspace card status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace card status activated successfully",
		Data:    "",
	})
}
