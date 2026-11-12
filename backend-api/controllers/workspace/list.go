package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListWorkspace(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := database.DB.First(&user).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed get user info", err)
		return
	}

	workspaces := []structs.WorkspaceResponse{}
	workspace_by_member := []uint{}

	database.DB.Model(&models.WorkspaceMember{UserId: user.Id}).Pluck("workspace_id", &workspace_by_member)

	for _, workspace_id := range workspace_by_member {
		workspace := structs.WorkspaceResponse{}
		database.DB.Model(&models.Workspace{Id: workspace_id}).First(&workspace)
		workspaces = append(workspaces, workspace)
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace retrieved successfully",
		Data:    workspaces,
	})
}
