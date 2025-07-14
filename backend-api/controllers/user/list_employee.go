package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListEmployee(c *gin.Context) {
	var user_id = c.GetInt("user_id")
	var user = models.User{Id: uint(user_id)}

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

	var employees = []models.User{}

	if err := database.DB.Where("role = ? AND is_deleted = ?", "employee", false).Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to get user info",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var list_employee = []structs.UserResponse{}

	for _, employee := range employees {
		list_employee = append(list_employee, structs.UserResponse{
			Id:        employee.Id,
			Name:      employee.Name,
			Username:  employee.Username,
			Email:     employee.Email,
			Role:      employee.Role,
			CreatedAt: employee.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: employee.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success to retreive employee",
		Data:    list_employee,
	})
}
