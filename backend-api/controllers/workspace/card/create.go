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
	workspaceId, _ := strconv.Atoi(id)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("id = ?", workspaceId).First(&models.Workspace{}).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Workspace not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("id = ? AND workspace_id = ? AND is_archived = ?", req.StatusId, workspaceId, false).First(&models.WorkspaceCardStatus{}).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Status card not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	result := database.DB.Model(&models.WorkspaceCard{}).Where("status_id = ?", req.StatusId)

	var card = models.WorkspaceCard{
		WorkspaceId: uint(workspaceId),
		StatusId:    uint(req.StatusId),
		Title:       req.Title,
		Description: req.Description,
		Order:       int(result.RowsAffected),
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
