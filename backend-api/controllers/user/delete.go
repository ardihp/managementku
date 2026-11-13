package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	user_id := c.Param("id")
	user := models.User{}

	if err := database.DB.First(&user, user_id).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "User not found", err)
		return
	}

	user.IsDeleted = true

	if err := database.DB.Save(&user).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to delete user", nil)
}
