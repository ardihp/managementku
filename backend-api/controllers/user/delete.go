package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	user_id := c.Param("id")
	user := models.User{}

	if err := database.DB.First(&user, user_id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user.IsDeleted = true

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to delete user",
	})

}
