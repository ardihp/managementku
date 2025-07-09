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

func DeleteWorkspace(c *gin.Context) {
	var id = c.Param("id")
	parseId, _ := strconv.Atoi(id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.First(&models.Workspace{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		workspace := models.Workspace{}
		trx.First(&workspace, parseId)

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "general",
			UserRole:    "manager",
			Action:      "delete-workspace",
			Description: fmt.Sprintf("User %s has deleted workspace %s", user.Username, workspace.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace",
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
					Message: "Failed to delete workspace",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}
		}

		if err := trx.Where("workspace_id = ?", id).Delete(&models.WorkspaceMember{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("workspace_id = ?", id).Delete(&models.WorkspaceTag{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("workspace_id = ?", id).Delete(&models.WorkspaceCard{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("workspace_id = ?", id).Delete(&models.WorkspaceCardStatus{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Delete(&models.Workspace{}, parseId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace deleted successfully",
		Data:    "",
	})
}
