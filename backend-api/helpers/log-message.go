package helpers

import (
	"kanban-board/backend-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, structs.ErrorResponse{
		Success: false,
		Message: message,
		Errors:  TranslateErrorMessage(err),
	})
}

func ValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
		Success: false,
		Message: "Validation Error",
		Errors:  TranslateErrorMessage(err),
	})
}
