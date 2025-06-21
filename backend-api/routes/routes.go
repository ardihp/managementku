package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"
	userControllers "kanban-board/backend-api/controllers/user"
	workspaceControllers "kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"
	workspaceStatusControllers "kanban-board/backend-api/controllers/workspace/status"
	workspaceTypeControllers "kanban-board/backend-api/controllers/workspace/type"

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

	// Authentication
	router.POST("/api/register", authControllers.RegisterUser)
	router.POST("/api/login", authControllers.LoginUser)

	// User
	router.GET("/api/employee", middlewares.AuthMiddleware(), userControllers.ListEmployee)

	// Workspace
	router.GET("/api/workspace", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspace)
	router.GET("/api/workspace/:id", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspaceByID)
	router.POST("/api/workspace/create", middlewares.AuthMiddleware(), workspaceControllers.CreateWorkspace)
	router.PUT("/api/workspace/update/:id", middlewares.AuthMiddleware(), workspaceControllers.UpdateWorkspace)
	router.PUT("/api/workspace/archive/:id", middlewares.AuthMiddleware(), workspaceControllers.ArchieveWorkspace)
	router.PUT("/api/workspace/unarchive/:id", middlewares.AuthMiddleware(), workspaceControllers.UnArchieveWorkspace)
	router.POST("/api/workspace/assign", middlewares.AuthMiddleware(), workspaceControllers.AssignMemberToWorkspace)
	router.GET("/api/workspace/archived", middlewares.AuthMiddleware(), workspaceControllers.ListArchivedWorkspace)
	router.DELETE("/api/workspace/delete/:id", middlewares.AuthMiddleware(), workspaceControllers.DeleteWorkspace)
	router.POST("/api/workspace/update-order", middlewares.AuthMiddleware(), workspaceControllers.UpdateWorkspaceOrder)

	// Workspace Status
	router.GET("/api/workspace/status", middlewares.AuthMiddleware(), workspaceStatusControllers.ListWorkspaceStatus)
	router.POST("/api/workspace/status/create", middlewares.AuthMiddleware(), workspaceStatusControllers.CreateWorkspaceStatus)

	// Workspace Type
	router.GET("/api/workspace/type", middlewares.AuthMiddleware(), workspaceTypeControllers.ListWorkspaceType)
	router.POST("/api/workspace/type/create", middlewares.AuthMiddleware(), workspaceTypeControllers.CreateWorkspaceType)

	// Workspace Card
	router.GET("/api/workspace/:id/card", middlewares.AuthMiddleware(), workspaceCardControllers.ListWorkspaceCard)
	router.POST("/api/workspace/:id/card/create", middlewares.AuthMiddleware(), workspaceCardControllers.CreateWorkspaceCard)
	router.POST("/api/workspace/:id/card/status-order", middlewares.AuthMiddleware(), workspaceCardControllers.UpdateWorkspaceCardStatusOrder)
	router.POST("/api/workspace/:id/card/update-order", middlewares.AuthMiddleware(), workspaceCardControllers.UpdateWorkspaceCardOrder)

	return router
}
