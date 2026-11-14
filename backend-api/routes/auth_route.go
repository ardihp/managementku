package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", authControllers.RegisterUser)
		auth.POST("/sign-in", authControllers.LoginUser)
	}
}
