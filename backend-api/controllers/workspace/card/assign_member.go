package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AssignMember(c *gin.Context) {
	workspace_id := c.Param("id")
	parsed_workspace_id, _ := strconv.Atoi(workspace_id)
	req := structs.WorkspaceCardAssignMember{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	manager_id := c.GetInt("user_id")
	manager := models.User{}
	database.DB.First(&manager, manager_id)

	workspace_card := models.WorkspaceCard{}
	members := []structs.MemberDetail{}

	user_log := models.HistoryLog{
		UserId:   uint(manager_id),
		Type:     "workspace",
		UserRole: "user",
		Action:   "assign-member-card",
	}

	manager_log := models.HistoryLog{
		UserId:   uint(manager_id),
		Type:     "workspace",
		UserRole: "admin",
		Action:   "assign-member-card",
	}

	workspace_log := models.WorkspaceHistoryLog{
		WorkspaceId: uint(parsed_workspace_id),
	}

	err := database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.Preload(clause.Associations).First(&workspace_card, req.CardId).Error; err != nil {
			return err
		}

		existing_members := []int{}
		trx.Model(&models.WorkspaceCardMember{}).Where("card_id = ?", req.CardId).Pluck("user_id", &existing_members)

		for _, member_id := range existing_members {
			user := models.User{}
			trx.First(&user, member_id)

			if !slices.Contains(req.UserId, int(member_id)) && member_id != manager_id {
				user_log.Description = fmt.Sprintf("User %s has remove you from card %s", manager.Username, workspace_card.Title)
				manager_log.Description = fmt.Sprintf("User %s remove %s from card %s", manager.Username, user.Username, workspace_card.Title)

				logs := []models.HistoryLog{user_log, manager_log}

				if err := trx.Create(&logs).Error; err != nil {
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryId: logs[0].ID, UserId: user.Id},
					{HistoryId: logs[1].ID, UserId: uint(manager_id)},
				}

				workspace_log.HistoryId = logs[1].ID

				if err := trx.Create(&workspace_log).Error; err != nil {
					return err
				}

				if err := trx.Create(&user_history).Error; err != nil {
					return err
				}
			}
		}

		assigned_user := []models.WorkspaceCardMember{}

		for key := range req.UserId {
			user := models.User{}
			exist_member := models.WorkspaceCardMember{}

			if err := trx.First(&user, req.UserId[key]).Error; err != nil {
				return err
			}

			trx.Where("card_id = ? AND user_id = ?", req.CardId, user.Id).First(&exist_member)

			assigned_user = append(assigned_user, models.WorkspaceCardMember{
				WorkspaceCardId: uint(req.CardId),
				UserId:          user.Id,
			})

			members = append(members, structs.MemberDetail{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     "user",
			})

			if user.Id != uint(manager_id) && !slices.Contains(existing_members, int(user.Id)) {
				user_log.Description = fmt.Sprintf("User %s has added you to card %s", manager.Username, workspace_card.Title)
				manager_log.Description = fmt.Sprintf("User %s added %s to card %s", manager.Username, user.Username, workspace_card.Title)

				logs := []models.HistoryLog{user_log, manager_log}

				if err := trx.Create(&logs).Error; err != nil {
					return err
				}

				user_history := []models.UserHistoryLog{
					{HistoryId: logs[0].ID, UserId: user.Id},
					{HistoryId: logs[1].ID, UserId: uint(manager_id)},
				}

				workspace_log.HistoryId = logs[1].ID

				if err := trx.Create(&workspace_log).Error; err != nil {
					return err
				}

				if err := trx.Create(&user_history).Error; err != nil {
					return err
				}
			}
		}

		if err := trx.Where("card_id = ?", req.CardId).Delete(&models.WorkspaceCardMember{}).Error; err != nil {
			return err
		}

		if len(assigned_user) >= 1 {
			if err := trx.Create(&assigned_user).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to assign user", err)
		return
	}

	helpers.SendSuccess(c, http.StatusCreated, "Update workspace card member successfully", "")
}
