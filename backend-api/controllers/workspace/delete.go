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
	workspace_id := c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)

	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := database.DB.First(&models.Workspace{}, parsed_workspace_id).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Workspace not found", err)
		return
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		workspace := models.Workspace{}
		trx.First(&workspace, parsed_workspace_id)

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "general",
			UserRole:    "manager",
			Action:      "delete-workspace",
			Description: fmt.Sprintf("User %s has deleted workspace %s", user.Username, workspace.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
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
				return err
			}
		}

		if err := trx.Where("workspace_id = ?", parsed_workspace_id).Delete(&models.WorkspaceMember{}).Error; err != nil {
			return err
		}

		if err := trx.Where("workspace_id = ?", parsed_workspace_id).Delete(&models.WorkspaceCard{}).Error; err != nil {
			return err
		}

		if err := trx.Where("workspace_id = ?", parsed_workspace_id).Delete(&models.WorkspaceCardStatus{}).Error; err != nil {
			return err
		}

		if err := trx.Delete(&models.Workspace{}, parsed_workspace_id).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to delete workspace", err)
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace deleted successfully",
		Data:    "",
	})
}
