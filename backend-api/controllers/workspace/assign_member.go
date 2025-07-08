package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AssignMemberToWorkspace(c *gin.Context) {
	var req = structs.WorkspaceAssignMember{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	manager_id := c.GetInt("user_id")
	manager := models.User{}
	database.DB.First(&manager, manager_id)

	workspace := models.Workspace{}
	members := []structs.Member{}
	workspace_tags := []string{}

	user_log := models.HistoryLog{
		UserId:   uint(manager_id),
		Type:     "general",
		UserRole: "employee",
		Action:   "assign-member",
	}

	manager_log := models.HistoryLog{
		UserId:   uint(manager_id),
		Type:     "general",
		UserRole: "manager",
		Action:   "assign-member",
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Preload(clause.Associations).First(&workspace, req.WorkspaceId).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		trx.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", req.WorkspaceId).Pluck("tag", &workspace_tags)

		existing_members := []int{}
		trx.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", req.WorkspaceId).Pluck("user_id", &existing_members)

		for _, member_id := range existing_members {
			user := models.User{}
			trx.First(&user, member_id)

			if !slices.Contains(req.UserId, int(member_id)) && member_id != manager_id {
				user_log.Description = fmt.Sprintf("User %s has remove you from workspace %s", manager.Username, workspace.Name)
				manager_log.Description = fmt.Sprintf("User %s remove %s from workspace %s", manager.Username, user.Username, workspace.Name)

				logs := []models.HistoryLog{user_log, manager_log}

				if err := trx.Create(&logs).Error; err != nil {
					c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
						Success: false,
						Message: "Failed to assign user",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryId: logs[0].ID, UserId: user.Id},
					{HistoryId: logs[1].ID, UserId: uint(manager_id)},
				}

				if err := trx.Create(&user_history).Error; err != nil {
					c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
						Success: false,
						Message: "Failed to update workspace order",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}
			}
		}

		assigned_user := []models.WorkspaceMember{}

		for key := range req.UserId {
			user := models.User{}
			exist_member := models.WorkspaceMember{}

			if err := trx.First(&user, req.UserId[key]).Error; err != nil {
				c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
					Success: false,
					Message: "User Not Found",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return err
			}

			trx.Where("workspace_id = ? AND user_id = ?", req.WorkspaceId, user.Id).First(&exist_member)

			role := ""

			if exist_member.Role != "" {
				role = exist_member.Role
			} else {
				role = "member"
			}

			assigned_user = append(assigned_user, models.WorkspaceMember{
				WorkspaceId: uint(req.WorkspaceId),
				UserId:      user.Id,
				Role:        role,
			})

			members = append(members, structs.Member{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     role,
			})

			if user.Id != uint(manager_id) && !slices.Contains(existing_members, int(user.Id)) {
				user_log.Description = fmt.Sprintf("User %s has added you to workspace %s", manager.Username, workspace.Name)
				manager_log.Description = fmt.Sprintf("User %s added %s to workspace %s", manager.Username, user.Username, workspace.Name)

				logs := []models.HistoryLog{user_log, manager_log}

				if err := trx.Create(&logs).Error; err != nil {
					c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
						Success: false,
						Message: "Failed to assign user",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryId: logs[0].ID, UserId: user.Id},
					{HistoryId: logs[1].ID, UserId: uint(manager_id)},
				}

				if err := trx.Create(&user_history).Error; err != nil {
					c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
						Success: false,
						Message: "Failed to update workspace order",
						Errors:  helpers.TranslateErrorMessage(err),
					})
					return err
				}
			}
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
