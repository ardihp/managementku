package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"
	workspaceControllers "kanban-board/backend-api/controllers/workspace"
	workspaceRoleControllers "kanban-board/backend-api/controllers/workspace/role"
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

	// Workspace
	router.GET("/api/workspace", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspace)
	router.GET("/api/workspace/:id", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspaceByID)
	router.POST("/api/workspace/create", middlewares.AuthMiddleware(), workspaceControllers.CreateWorkspace)
	router.PUT("/api/workspace/update/:id", middlewares.AuthMiddleware(), workspaceControllers.UpdateWorkspace)
	router.POST("/api/workspace/assign", middlewares.AuthMiddleware(), workspaceControllers.AssignMemberToWorkspace)

	// Workspace Type
	router.POST("/api/workspace/type/create", middlewares.AuthMiddleware(), workspaceTypeControllers.CreateWorkspaceType)

	// Workspace Status
	router.POST("/api/workspace/update/status", middlewares.AuthMiddleware(), workspaceControllers.UpdateStatusWorkspace)

	// Workspace Role
	router.POST("/api/workspace/role/create", middlewares.AuthMiddleware(), workspaceRoleControllers.CreateWorkspaceRole)

	return router
}
