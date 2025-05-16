package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"
	workspaceControllers "kanban-board/backend-api/controllers/workspace"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/api/register", authControllers.RegisterUser)
	router.POST("/api/login", authControllers.LoginUser)
	router.POST("/api/workspace/create", workspaceControllers.CreateWorkspace)

	return router
}
