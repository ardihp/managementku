package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AssignMemberToWorkspace(c *gin.Context) {
	var req = structs.WorkspaceAssignMember{}
	var workspace = models.Workspace{Id: uint(req.WorkspaceId)}
	var members = []structs.Member{}
	var workspace_tags = []string{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Preload(clause.Associations).First(&workspace).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", req.WorkspaceId).Pluck("tag", &workspace_tags)

	database.DB.Transaction(func(trx *gorm.DB) error {
		var assigned_user = []models.WorkspaceMember{}

		for key := range req.UserId {
			var user = models.User{}
			var existing_member = models.WorkspaceMember{}

			if err := trx.First(&user, req.UserId[key]).Error; err != nil {
				c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
					Success: false,
					Message: "User Not Found",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}

			trx.Where("workspace_id = ? AND user_id = ?", req.WorkspaceId, req.UserId[key]).First(&existing_member)

			var role = ""

			if existing_member.Role != "" {
				role = existing_member.Role
			} else {
				role = "member"
			}

			assigned_user = append(assigned_user, models.WorkspaceMember{
				WorkspaceId: uint(req.WorkspaceId),
				UserId:      uint(req.UserId[key]),
				Role:        role,
			})

			members = append(members, structs.Member{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     "member",
			})
		}

		if err := trx.Where("workspace_id = ?", req.WorkspaceId).Delete(&models.WorkspaceMember{}).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Failed to assign user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Create(&assigned_user).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Failed to assign user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User assigned successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Name:  workspace.Type.Name,
				Color: workspace.Type.Color,
			},
			Members:   members,
			Tags:      workspace_tags,
			CreatedAt: workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
