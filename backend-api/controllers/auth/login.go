package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	if err := database.DB.Where("username = ? AND is_deleted = ?", req.Username, false).First(&user).Error; err != nil {
		helpers.SendError(c, http.StatusUnauthorized, "User Not Found", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		helpers.SendError(c, http.StatusUnauthorized, "Invalid Password", err)
		return
	}

	token := helpers.GenerateToken(user.Username, int(user.Id))

	helpers.SendSuccess(c, http.StatusOK, "Login Success", structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		Token:     &token,
	})
}
