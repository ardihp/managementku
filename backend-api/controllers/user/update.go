package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	req := structs.UserUpdateRequest{}
	user_id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{}

	if err := database.DB.First(&user, user_id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.Role = req.Role

	if req.Password != "" {
		user.Password = helpers.HashPassword(req.Password)
	}

	if err := database.DB.Save(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to save user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to save user",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
