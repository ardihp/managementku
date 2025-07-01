package controllers

import (
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteWorkspaceStatus(c *gin.Context) {
	id := c.Param("id")

	database.DB.Transaction(func(trx *gorm.DB) error {
		if err := trx.First(&models.WorkspaceStatus{}, id).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Workspace status not found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Where("status_id = ?", id).Delete(&models.Workspace{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		if err := trx.Delete(&models.WorkspaceStatus{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to delete workspace status",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return err
		}

		return nil
	})

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success delete workspace status",
	})
}
