package user

import (
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserController) UpdateUser(ctx *gin.Context) {
	req := structs.UserUpdateRequest{}
	user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(ctx, err)
		return
	}

	user := models.User{}

	if err := c.DB.First(&user, user_id).Error; err != nil {
		helpers.SendError(ctx, http.StatusUnprocessableEntity, "User not found", err)
		return
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email

	if req.Password != "" {
		user.Password = helpers.HashPassword(req.Password)
	}

	if err := c.DB.Save(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			helpers.SendError(ctx, http.StatusConflict, "Duplicate entry error", err)
		} else {
			helpers.SendError(ctx, http.StatusInternalServerError, "Failed to save user", err)
		}
		return
	}

	helpers.SendSuccess(ctx, http.StatusOK, "Success to save user", structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
