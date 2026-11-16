package user

import (
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserController) DeleteUser(ctx *gin.Context) {
	user_id := ctx.Param("id")
	user := models.User{}

	if err := c.DB.First(&user, user_id).Error; err != nil {
		helpers.SendError(ctx, http.StatusUnprocessableEntity, "User not found", err)
		return
	}

	user.IsDeleted = true

	if err := c.DB.Save(&user).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	helpers.SendSuccess(ctx, http.StatusOK, "Success to delete user", nil)
}
