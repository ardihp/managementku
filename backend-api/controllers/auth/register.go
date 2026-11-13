package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			helpers.SendError(c, http.StatusConflict, "Duplicate entry error", err)
		} else {
			helpers.SendError(c, http.StatusInternalServerError, "Failed to create user", err)
		}
		return
	}

	token := helpers.GenerateToken(user.Username, int(user.Id))

	helpers.SendSuccess(c, http.StatusCreated, "User created successfully", structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		Token:     &token,
	})
}
