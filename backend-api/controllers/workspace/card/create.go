package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateWorkspaceCard(c *gin.Context) {
	var req = structs.WorkspaceCreateCardRequest{}
	var id = c.Param("id")
	parseId, _ := strconv.Atoi(id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var card = models.WorkspaceCard{
		WorkspaceId: parseId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	if err := database.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create workspace card",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Workspace card created successfully",
		Data: structs.WorkspaceCardResponse{
			Id:          card.Id,
			Title:       card.Title,
			Description: card.Description,
			Members:     []structs.Member{},
		},
	})
}
