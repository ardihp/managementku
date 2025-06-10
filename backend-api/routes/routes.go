package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"
	workspaceControllers "kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"
	workspaceRoleControllers "kanban-board/backend-api/controllers/workspace/role"
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

	// Workspace
	router.GET("/api/workspace", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspace)
	router.GET("/api/workspace/:id", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspaceByID)
	router.POST("/api/workspace/create", middlewares.AuthMiddleware(), workspaceControllers.CreateWorkspace)
	router.PUT("/api/workspace/update/:id", middlewares.AuthMiddleware(), workspaceControllers.UpdateWorkspace)
	router.PUT("/api/workspace/archive/:id", middlewares.AuthMiddleware(), workspaceControllers.ArchieveWorkspace)
	router.PUT("/api/workspace/unarchive/:id", middlewares.AuthMiddleware(), workspaceControllers.UnArchieveWorkspace)
	router.POST("/api/workspace/assign", middlewares.AuthMiddleware(), workspaceControllers.AssignMemberToWorkspace)
	router.DELETE("/api/workspace/delete/:id", middlewares.AuthMiddleware(), workspaceControllers.DeleteWorkspace)

	// Workspace Status
	router.GET("/api/workspace/:id/status", middlewares.AuthMiddleware(), workspaceStatusControllers.ListWorkspaceStatus)
	router.POST("/api/workspace/update/status", middlewares.AuthMiddleware(), workspaceControllers.UpdateStatusWorkspace)

	// Workspace Type
	router.GET("/api/workspace/type", middlewares.AuthMiddleware(), workspaceTypeControllers.ListWorkspaceType)
	router.POST("/api/workspace/type/create", middlewares.AuthMiddleware(), workspaceTypeControllers.CreateWorkspaceType)

	// Workspace Role
	router.GET("/api/workspace/role", middlewares.AuthMiddleware(), workspaceRoleControllers.ListWorkspaceRole)
	router.POST("/api/workspace/role/create", middlewares.AuthMiddleware(), workspaceRoleControllers.CreateWorkspaceRole)

	// Workspace Card
	router.POST("/api/workspace/:id/card/create", middlewares.AuthMiddleware(), workspaceCardControllers.CreateWorkspaceCard)

	return router
}
