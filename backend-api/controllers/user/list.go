package user

import (
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserController) ListUser(ctx *gin.Context) {
	user_id := ctx.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := c.DB.First(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to get user info", err)
		return
	}

	users := []models.User{}

	if err := c.DB.Find(&users, "is_deleted = ?", false).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to retreive users", err)
		return
	}

	list_users := []structs.UserResponse{}

	for _, user := range users {
		list_users = append(list_users, structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			IsDeleted: user.IsDeleted,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendSuccess(ctx, http.StatusOK, "Success to retreive users", list_users)
}
