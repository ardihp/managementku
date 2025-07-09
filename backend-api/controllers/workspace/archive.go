package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArchieveWorkspace(c *gin.Context) {
	var workspace = models.Workspace{}
	var id = c.Param("id")

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if workspace.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace is already archived",
			Errors:  nil,
		})
		return
	}

	workspace.IsArchived = true

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "general",
			UserRole:    "manager",
			Action:      "archive-workspace",
			Description: fmt.Sprintf("User %s has archived workspace %s", user.Username, workspace.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to archive workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_members := []models.WorkspaceMember{}
		trx.Where("workspace_id = ?", workspace.Id).Find(&workspace_members)

		for _, member := range workspace_members {
			user_history := models.UserHistoryLog{
				HistoryId: history_log.ID,
				UserId:    member.UserId,
			}

			if err := trx.Create(&user_history).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to archive workspace",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace archived successfully",
		Data:    "",
	})
}

func UnArchieveWorkspace(c *gin.Context) {
	var workspace = models.Workspace{}
	var id = c.Param("id")

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if !workspace.IsArchived {
		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Workspace is already active",
			Errors:  nil,
		})
		return
	}

	workspace.IsArchived = false

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "general",
			UserRole:    "manager",
			Action:      "unarchive-workspace",
			Description: fmt.Sprintf("User %s has unarchived workspace %s", user.Username, workspace.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to unarchive workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		workspace_members := []models.WorkspaceMember{}
		trx.Where("workspace_id = ?", workspace.Id).Find(&workspace_members)

		for _, member := range workspace_members {
			user_history := models.UserHistoryLog{
				HistoryId: history_log.ID,
				UserId:    member.UserId,
			}

			if err := trx.Create(&user_history).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to unarchive workspace",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace activated successfully",
		Data:    "",
	})
}
