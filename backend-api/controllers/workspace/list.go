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

func ListWorkspace(c *gin.Context) {
	var user_id = c.GetInt("user_id")
	var user = models.User{Id: uint(user_id)}

	if err := database.DB.First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed get user info",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var workspace_status = []models.WorkspaceStatus{}
	database.DB.Where("is_archived = ?", false).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false}).Find(&workspace_status)

	var list_workspace = []structs.WorkspaceByStatusResponse{}

	for _, status := range workspace_status {
		var user_workspace = []int{}

		if user.Role == "employee" {
			var list_workspace = []models.WorkspaceMember{}
			database.DB.Model(&models.WorkspaceMember{}).Preload(clause.Associations).Where("user_id = ?", user_id).Order("created_at").Find(&list_workspace)

			for _, list := range list_workspace {
				workspace := &models.Workspace{}
				database.DB.First(&workspace, list.WorkspaceId)

				if !workspace.IsArchived {
					if list.Workspace.StatusId == status.Id {
						user_workspace = append(user_workspace, int(list.WorkspaceId))
					}
				}
			}
		} else {
			database.DB.Model(&models.Workspace{}).Where("is_archived = ? AND status_id = ?", false, status.Id).Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}}).Pluck("id", &user_workspace)
		}

		var list_workspace_by_status = []structs.WorkspaceResponse{}

		for _, workspace_id := range user_workspace {
			var members = []structs.Member{}
			var id_member = []int{}
			var workspace = models.Workspace{
				Id: uint(workspace_id),
			}
			var workspace_tags = []string{}

			database.DB.Model(&models.Workspace{}).Preload(clause.Associations).First(&workspace)
			database.DB.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace_id).Order("user_id").Pluck("user_id", &id_member)
			database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", workspace_id).Pluck("tag", &workspace_tags)

			for _, user_id := range id_member {
				var user = models.User{}
				var workspace_member = models.WorkspaceMember{}

				database.DB.First(&user, user_id)
				database.DB.Model(&models.WorkspaceMember{}).Preload(clause.Associations).Where("workspace_id = ? AND user_id = ?", workspace_id, user_id).Find(&workspace_member)

				members = append(members, structs.Member{
					Id:       user.Id,
					Name:     user.Name,
					Username: user.Username,
					Email:    user.Email,
					Role:     workspace_member.Role,
				})
			}

			list_workspace_by_status = append(list_workspace_by_status, structs.WorkspaceResponse{
				Id:          workspace.Id,
				Name:        workspace.Name,
				Description: workspace.Description,
				Type: structs.WorkspaceType{
					Id:    workspace.Type.Id,
					Name:  workspace.Type.Name,
					Color: workspace.Type.Color,
				},
				Status: structs.WorkspaceStatus{
					Id:         workspace.Status.Id,
					Name:       workspace.Status.Name,
					Color:      workspace.Status.Color,
					IsArchived: workspace.Status.IsArchived,
				},
				Members:    members,
				Tags:       workspace_tags,
				IsArchived: workspace.IsArchived,
				CreatedAt:  workspace.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:  workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		list_workspace = append(list_workspace, structs.WorkspaceByStatusResponse{
			Id:         status.Id,
			Name:       status.Name,
			Color:      status.Color,
			IsArchived: status.IsArchived,
			Workspaces: list_workspace_by_status,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace retrieved successfully",
		Data:    list_workspace,
	})
}

func ListWorkspaceByID(c *gin.Context) {
	var workspace = models.Workspace{}
	var workspace_id = c.Param("id")
	var workspace_tags = []string{}

	if err := database.DB.Model(&models.Workspace{}).Preload(clause.Associations).First(&workspace, workspace_id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var members = []structs.Member{}
	var id_member = []int{}

	database.DB.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member)
	database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", workspace_id).Pluck("tag", &workspace_tags)

	for _, user_id := range id_member {
		var user = models.User{}
		var workspace_member = models.WorkspaceMember{}

		database.DB.First(&user, user_id)
		database.DB.Model(&models.WorkspaceMember{}).Preload(clause.Associations).Where("workspace_id = ? AND user_id = ?", workspace_id, user_id).First(&workspace_member)

		members = append(members, structs.Member{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role:     workspace_member.Role,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Workspace retrieved successfully",
		Data: structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Id:    workspace.Type.Id,
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

func ListArchivedWorkspace(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := database.DB.First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed get user info",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if user.Role != "manager" {
		c.JSON(http.StatusForbidden, structs.ErrorResponse{
			Success: false,
			Message: "You are not allowed to access this resource",
			Errors:  map[string]string{},
		})
		return
	}

	var workspaces = []models.Workspace{}

	if err := database.DB.Model(&models.Workspace{}).Preload(clause.Associations).Where("is_archived = ?", true).Order("updated_at").Find(&workspaces).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive archived workspace",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var list_workspace = []structs.WorkspaceResponse{}

	for _, workspace := range workspaces {
		var id_member = []int{}
		var workspace_tags = []string{}

		database.DB.Model(&models.WorkspaceMember{}).Where("workspace_id = ?", workspace.Id).Order("user_id").Pluck("user_id", &id_member)
		database.DB.Model(&models.WorkspaceTag{}).Where("workspace_id = ?", workspace.Id).Pluck("tag", &workspace_tags)

		var members = []structs.Member{}

		for _, user_id := range id_member {
			var user = models.User{}
			var workspace_member = models.WorkspaceMember{}

			database.DB.First(&user, user_id)
			database.DB.Model(&models.WorkspaceMember{}).Preload(clause.Associations).Where("workspace_id = ? AND user_id = ?", workspace.Id, user_id).Find(&workspace_member)

			members = append(members, structs.Member{
				Id:       user.Id,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Role:     workspace_member.Role,
			})
		}

		list_workspace = append(list_workspace, structs.WorkspaceResponse{
			Id:          workspace.Id,
			Name:        workspace.Name,
			Description: workspace.Description,
			Type: structs.WorkspaceType{
				Id:    workspace.Type.Id,
				Name:  workspace.Type.Name,
				Color: workspace.Type.Color,
			},
			Status: structs.WorkspaceStatus{
				Id:    workspace.Status.Id,
				Name:  workspace.Status.Name,
				Color: workspace.Status.Color,
			},
			Tags:       workspace_tags,
			Members:    members,
			IsArchived: workspace.IsArchived,
			CreatedAt:  workspace.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  workspace.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Archived workspace retrieved successfully",
		Data:    list_workspace,
	})
}
