package controllers

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWorkspaceOrder(c *gin.Context) {
	var req = structs.WorkspaceUpdateOrderRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	for _, order := range req.OrderStatus {
		for index, workspace := range order.WorkspaceId {
			fmt.Println(order.WorkspaceStatusId, workspace)

			var new_workspace = models.Workspace{Id: uint(workspace)}
			database.DB.First(&new_workspace)

			new_workspace.WorkspaceStatus.Id = uint(order.WorkspaceStatusId)
			new_workspace.Order = index

			if err := database.DB.Save(&new_workspace).Error; err != nil {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to update workspace order",
					Errors:  helpers.TranslateErrorMessage(err),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success update workspace order",
		Data:    "",
	})
}
