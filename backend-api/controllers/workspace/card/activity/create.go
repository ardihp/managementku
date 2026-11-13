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
		helpers.ValidationError(c, err)
		return
	}

	if err := database.DB.First(&models.WorkspaceCard{}, card_id).Error; err != nil {
		helpers.SendError(c, http.StatusUnprocessableEntity, "Workspace card is not found", err)
		return
	}

	workspace_card_activity := models.WorkspaceCardActivity{
		WorkspaceCardId: uint(parsed_card_id),
		UserId:          uint(user_id),
		Type:            "comment",
		Message:         req.Message,
	}

	if err := database.DB.Create(&workspace_card_activity).Error; err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "Failed to create workspace activity", err)
		return
	}

	helpers.SendSuccess(c, http.StatusOK, "Success create workspace card activity", "")
}
