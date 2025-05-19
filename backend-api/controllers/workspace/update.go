package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWorkspace(c *gin.Context) {
	var req = structs.WorkspaceUpdateRequest{}
	var workspace = models.Workspace{}
	var id = c.Param("id")
	// parseId, _ := strconv.Atoi(id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.First(&workspace, id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace.Name = req.Name
	workspace.Description = req.Description
	workspace.Category = req.Category

	var members = []structs.Member{}
	var id_member = []int{}

	var workspace_member = []models.WorkspaceMember{}
	database.DB.Where("workspace_id = ?", workspace.Id).Order("user_id").Find(&workspace_member)

	if len(req.Members) >= 1 {
		for _, user_id := range req.Members {
			if err := database.DB.First(&models.User{}, user_id).Error; err != nil {
				c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
					Success: false,
					Message: "User not found",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}
		}

		// if len(req.Members) < len(workspace_member) {
		// 	var id_member_exist = []int{}
		// 	database.DB.Model(&workspace_member).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member_exist)

		// 	// database.DB.Where("workspace_id = ?", workspace.Id).Delete(&workspace_member, models.WorkspaceMember{UserId: data.UserId})
		// } else {
		// 	var id_member_exist = []int{}
		// 	database.DB.Model(&workspace_member).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member_exist)

		// 	for _, user_id := range req.Members {

		// 		if !slices.Contains(id_member_exist, user_id) {
		// 			database.DB.Create(&models.WorkspaceMember{WorkspaceId: uint(parseId), UserId: uint(user_id)})
		// 		}
		// 		id_member = append(id_member, user_id)
		// 	}
		// }
	} else {
		database.DB.Model(&workspace_member).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member)
	}

	for _, id := range id_member {
		var user = models.User{}

		database.DB.First(&user, id)
		members = append(members, structs.Member{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	if err := database.DB.Save(&workspace).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace updated successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Category:    workspace.Category,
			Members:     members,
			CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
