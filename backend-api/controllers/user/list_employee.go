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
		helpers.SendError(c, http.StatusInternalServerError, "Failed to get user info", err)
		return
	}

	var employees = []models.User{}

	if err := database.DB.Where("is_deleted = ?", false).Find(&employees).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to get user info", err)
		return
	}

	var list_employee = []structs.UserResponse{}

	for _, employee := range employees {
		list_employee = append(list_employee, structs.UserResponse{
			Id:        employee.Id,
			Name:      employee.Name,
			Username:  employee.Username,
			Email:     employee.Email,
			CreatedAt: employee.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: employee.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendSuccess(c, http.StatusOK, "Success to retreive employee", list_employee)
}
