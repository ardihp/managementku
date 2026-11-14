package routes

import (
	socketControllers "kanban-board/backend-api/controllers/socket"

	"kanban-board/backend-api/middlewares"

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

	//Web Socket
	router.GET("/ws", socketControllers.InitSocket)

	api := router.Group("/api")
	{
		authRoutes(api)

		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			userRoute(protected)
			workspaceRoute(protected)
			workspaceCardRoute(protected)
		}
	}

	return router
}
