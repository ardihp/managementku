package routes

import (
	"kanban-board/backend-api/controllers"
	socketControllers "kanban-board/backend-api/controllers/socket"
	"kanban-board/backend-api/database"

	"kanban-board/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db := database.DB
	ctrls := controllers.InitControllers(db)

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
			userRoute(protected, ctrls.User)
			workspaceRoute(protected, ctrls.Workspace)
			workspaceCardRoute(protected)
		}
	}

	return router
}
