package workspace

import (
	"fmt"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *WorkspaceController) AssignMember(ctx *gin.Context) {
	req := structs.WorkspaceAssignMember{}
	user_id := ctx.GetString("user_id")
	workspace_id := ctx.Param("id")

	user := models.User{}
	c.DB.First(&user, user_id)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(ctx, err)
		return
	}

	workspace := models.Workspace{}
	members := []structs.MemberDetail{}

	member_log := models.HistoryLog{
		UserID:   user_id,
		Type:     "general",
		UserRole: "user",
		Action:   "assign-member",
	}

	user_log := models.HistoryLog{
		UserID:   user_id,
		Type:     "general",
		UserRole: "admin",
		Action:   "assign-member",
	}

	err := c.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.First(&workspace, workspace_id).Error; err != nil {
			return err
		}

		existing_members := []string{}
		trx.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace_id).Pluck("user_id", &existing_members)

		for _, member_id := range existing_members {
			user := models.User{}
			trx.First(&user, member_id)

			if !slices.Contains(req.UserID, member_id) && member_id != user_id {
				member_log.Description = fmt.Sprintf("User %s has remove you from workspace %s", user.Username, workspace.Name)
				user_log.Description = fmt.Sprintf("User %s remove %s from workspace %s", user.Username, user.Username, workspace.Name)

				logs := []models.HistoryLog{member_log, user_log}

				if err := trx.Create(&logs).Error; err != nil {
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryID: logs[0].ID, UserID: user.Id},
					{HistoryID: logs[1].ID, UserID: user_id},
				}

				if err := trx.Create(&user_history).Error; err != nil {
					return err
				}
			}
		}

		assigned_user := []models.WorkspaceMember{}

		for key := range req.UserID {
			user := models.User{}
			exist_member := models.WorkspaceMember{}

			if err := trx.First(&user, req.UserID[key]).Error; err != nil {
				return err
			}

			trx.Where("workspace_id = ? AND user_id = ?", workspace_id, user.Id).First(&exist_member)

			role := ""

			if exist_member.Role != "" {
				role = exist_member.Role
			} else {
				role = "user"
			}

			assigned_user = append(assigned_user, models.WorkspaceMember{
				WorkspaceID: workspace_id,
				UserID:      user.Id,
				Role:        role,
			})

			members = append(members, structs.MemberDetail{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     role,
			})

			if user.Id != user_id && !slices.Contains(existing_members, user.Id) {
				member_log.Description = fmt.Sprintf("User %s has added you to workspace %s", user.Username, workspace.Name)
				user_log.Description = fmt.Sprintf("User %s added %s to workspace %s", user.Username, user.Username, workspace.Name)

				logs := []models.HistoryLog{user_log, user_log}

				if err := trx.Create(&logs).Error; err != nil {
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryID: logs[0].ID, UserID: user.Id},
					{HistoryID: logs[1].ID, UserID: user_id},
				}

				if err := trx.Create(&user_history).Error; err != nil {
					return err
				}
			}
		}

		if err := trx.Where("workspace_id = ?", workspace_id).Delete(&models.WorkspaceMember{}).Error; err != nil {
			return err
		}

		if err := trx.Create(&assigned_user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed assign member to workspace", err)
		return
	}

	helpers.SendSuccess(ctx, http.StatusCreated, "User assigned successfully", structs.WorkspaceResponse{
		Id:          workspace.Id,
		Name:        workspace.Name,
		Description: workspace.Description,
		Members:     members,
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
	})
}
