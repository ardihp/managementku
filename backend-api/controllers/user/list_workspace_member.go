package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ListWorkspaceMember(c *gin.Context) {
	workspace_id := c.Param("id")

	if err := database.DB.Preload(clause.Associations).First(&models.Workspace{}, workspace_id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_members := []models.WorkspaceMember{}

	if err := database.DB.Where("workspace_id = ? AND role = ?", workspace_id, "member").Preload(clause.Associations).Find(&workspace_members).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive workspace member",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	list_workspace_members := []structs.WorkspaceMember{}

	for _, member := range workspace_members {
		list_workspace_members = append(list_workspace_members, structs.WorkspaceMember{
			Id:       member.User.Id,
			Name:     member.User.Name,
			Username: member.User.Username,
			Email:    member.User.Email,
			Role:     member.User.Role,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive workspace member",
		Data:    list_workspace_members,
	})
}
