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

func UpdateWorkspace(c *gin.Context) {
	req := structs.WorkspaceUpdateRequest{}
	workspace_id := c.Param("id")
	user_id := c.GetInt("user_id")

	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	workspace := models.Workspace{}

	if err := database.DB.First(&workspace, workspace_id).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace not found", err)
		return
	}

	workspace.Name = req.Name
	workspace.Description = req.Description

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Save(&workspace).Error; err != nil {
			return err
		}

		history_log := models.HistoryLog{
			UserId:      user.Id,
			Type:        "general",
			UserRole:    "manager",
			Action:      "update-workspace",
			Description: fmt.Sprintf("User %s has updated workspace %s", user.Username, workspace.Name),
		}

		if err := trx.Create(&history_log).Error; err != nil {
			return err
		}

		user_history := models.UserHistoryLog{
			HistoryId: history_log.ID,
			UserId:    user.Id,
		}

		if err := trx.Create(&user_history).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to create workspace", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Workspace updated successfully", structs.WorkspaceResponse{
		Id:          workspace.Id,
		Name:        workspace.Name,
		Description: workspace.Description,
		Members:     []structs.WorkspaceMember{},
		CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
