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

func CreateWorkspaceCardActivity(c *gin.Context) {
	req := structs.WorkspaceCardCreateActivityRequst{}
	card_id := c.Param("card_id")
	parsed_card_id, _ := strconv.Atoi(card_id)

	user_id := c.GetInt("user_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.First(&models.WorkspaceCard{}, card_id).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Workspace card is not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	workspace_card_activity := models.WorkspaceCardActivity{
		WorkspaceCardId: uint(parsed_card_id),
		UserId:          uint(user_id),
		Type:            "comment",
		Message:         req.Message,
	}

	if err := database.DB.Create(&workspace_card_activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create workspace activity",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success create workspace card activity",
		Data:    "",
	})
}
