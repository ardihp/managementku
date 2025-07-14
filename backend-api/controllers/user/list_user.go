package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user := models.User{Id: uint(user_id)}

	if err := database.DB.First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to get user info",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if user.Role != "manager" {
		c.JSON(http.StatusForbidden, structs.ErrorResponse{
			Success: false,
			Message: "You don't have access",
			Errors:  map[string]string{},
		})
		return
	}

	users := []models.User{}

	if err := database.DB.Find(&users, "is_deleted = ?", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to retreive users",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	list_users := []structs.UserResponse{}

	for _, user := range users {
		list_users = append(list_users, structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			IsDeleted: user.IsDeleted,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive users",
		Data:    list_users,
	})
}
