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
	req := structs.WorkspaceAssignMember{}
	user_id := c.GetInt("user_id")
	user := models.User{}
	database.DB.First(&user, user_id)

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	workspace := models.Workspace{}
	members := []structs.WorkspaceMember{}

	member_log := models.HistoryLog{
		UserId:   uint(user_id),
		Type:     "general",
		UserRole: "employee",
		Action:   "assign-member",
	}

	user_log := models.HistoryLog{
		UserId:   uint(user_id),
		Type:     "general",
		UserRole: "manager",
		Action:   "assign-member",
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Preload(clause.Associations).First(&workspace, req.WorkspaceId).Error; err != nil {
			return err
		}

		existing_members := []int{}
		trx.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", req.WorkspaceId).Pluck("user_id", &existing_members)

		for _, member_id := range existing_members {
			user := models.User{}
			trx.First(&user, member_id)

			if !slices.Contains(req.UserId, int(member_id)) && member_id != user_id {
				member_log.Description = fmt.Sprintf("User %s has remove you from workspace %s", user.Username, workspace.Name)
				user_log.Description = fmt.Sprintf("User %s remove %s from workspace %s", user.Username, user.Username, workspace.Name)

				logs := []models.HistoryLog{member_log, user_log}

				if err := trx.Create(&logs).Error; err != nil {
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryId: logs[0].ID, UserId: user.Id},
					{HistoryId: logs[1].ID, UserId: uint(user_id)},
				}

				if err := trx.Create(&user_history).Error; err != nil {
					return err
				}
			}
		}

		assigned_user := []models.WorkspaceMember{}

		for key := range req.UserId {
			user := models.User{}
			exist_member := models.WorkspaceMember{}

			if err := trx.First(&user, req.UserId[key]).Error; err != nil {
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

			members = append(members, structs.WorkspaceMember{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     role,
			})

			if user.Id != uint(user_id) && !slices.Contains(existing_members, int(user.Id)) {
				member_log.Description = fmt.Sprintf("User %s has added you to workspace %s", user.Username, workspace.Name)
				user_log.Description = fmt.Sprintf("User %s added %s to workspace %s", user.Username, user.Username, workspace.Name)

				logs := []models.HistoryLog{user_log, user_log}

				if err := trx.Create(&logs).Error; err != nil {
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryId: logs[0].ID, UserId: user.Id},
					{HistoryId: logs[1].ID, UserId: uint(user_id)},
				}

				if err := trx.Create(&user_history).Error; err != nil {
					return err
				}
			}
		}

		if err := trx.Where("workspace_id = ?", req.WorkspaceId).Delete(&models.WorkspaceMember{}).Error; err != nil {
			return err
		}

		if err := trx.Create(&assigned_user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed assign member to workspace", err)
		return
	}

	helpers.SendSuccess(c, http.StatusCreated, "User assigned successfully", structs.WorkspaceResponse{
		Id:          workspace.Id,
		Name:        workspace.Name,
		Description: workspace.Description,
		Members:     members,
		CreatedAt:   workspace.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
