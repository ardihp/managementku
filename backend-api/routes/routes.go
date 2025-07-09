package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"
	historyControllers "kanban-board/backend-api/controllers/history"
	socketControllers "kanban-board/backend-api/controllers/socket"
	userControllers "kanban-board/backend-api/controllers/user"
	workspaceControllers "kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"
	workspaceCardStatusControllers "kanban-board/backend-api/controllers/workspace/card/status"
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

	//Web Socket
	router.GET("/ws", socketControllers.InitSocket)

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
	router.GET("/api/workspace/status/archived", middlewares.AuthMiddleware(), workspaceStatusControllers.ListArchivedWorkspaceStatus)
	router.POST("/api/workspace/status/create", middlewares.AuthMiddleware(), workspaceStatusControllers.CreateWorkspaceStatus)
	router.PUT("/api/workspace/status/update", middlewares.AuthMiddleware(), workspaceStatusControllers.UpdateWorkspaceStatus)
	router.PUT("/api/workspace/status/update-order", middlewares.AuthMiddleware(), workspaceStatusControllers.UpdateOrderWorkspaceStatus)
	router.PUT("/api/workspace/status/archive/:id", middlewares.AuthMiddleware(), workspaceStatusControllers.ArchieveWorkspaceStatus)
	router.PUT("/api/workspace/status/unarchive/:id", middlewares.AuthMiddleware(), workspaceStatusControllers.UnArchieveWorkspaceStatus)
	router.DELETE("/api/workspace/status/delete/:id", middlewares.AuthMiddleware(), workspaceStatusControllers.DeleteWorkspaceStatus)

	// Workspace Type
	router.GET("/api/workspace/type", middlewares.AuthMiddleware(), workspaceTypeControllers.ListWorkspaceType)
	router.POST("/api/workspace/type/create", middlewares.AuthMiddleware(), workspaceTypeControllers.CreateWorkspaceType)
	router.PUT("/api/workspace/type/update", middlewares.AuthMiddleware(), workspaceTypeControllers.UpdateWorkspaceType)
	router.DELETE("/api/workspace/type/delete/:id", middlewares.AuthMiddleware(), workspaceTypeControllers.DeleteWorkspaceType)

	// Workspace Card
	router.GET("/api/workspace/:id/card", middlewares.AuthMiddleware(), workspaceCardControllers.ListWorkspaceCard)
	router.GET("/api/workspace/:id/card/archived", middlewares.AuthMiddleware(), workspaceCardControllers.ListArchivedWorkspaceCard)
	router.POST("/api/workspace/:id/card/create", middlewares.AuthMiddleware(), workspaceCardControllers.CreateWorkspaceCard)
	router.PUT("/api/workspace/:id/card/update/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.UpdateWorkspaceCard)
	router.PUT("/api/workspace/:id/card/archive/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.ArchieveWorkspaceCard)
	router.PUT("/api/workspace/:id/card/unarchive/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.UnarchieveWorkspaceCard)
	router.DELETE("/api/workspace/:id/card/delete/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.DeleteWorkspaceCard)
	router.POST("/api/workspace/:id/card/status-order", middlewares.AuthMiddleware(), workspaceCardControllers.UpdateWorkspaceCardStatusOrder)
	router.POST("/api/workspace/:id/card/update-order", middlewares.AuthMiddleware(), workspaceCardControllers.UpdateWorkspaceCardOrder)

	// Workspace Card Status
	router.GET("/api/workspace/:id/status", middlewares.AuthMiddleware(), workspaceCardStatusControllers.ListWorkspaceCardStatus)
	router.GET("/api/workspace/:id/status/archived", middlewares.AuthMiddleware(), workspaceCardStatusControllers.ListArchivedWorkspaceCardStatus)
	router.POST("/api/workspace/:id/status/create", middlewares.AuthMiddleware(), workspaceCardStatusControllers.CreateWorkspaceCardStatus)
	router.PUT("/api/workspace/:id/status/update/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.UpdateWorkspaceCardStatus)
	router.PUT("/api/workspace/:id/status/archive/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.ArchieveWorkspaceCardStatus)
	router.PUT("/api/workspace/:id/status/unarchive/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.UnArchieveWorkspaceCardStatus)
	router.DELETE("/api/workspace/:id/status/delete/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.DeleteWorkspaceCardStatus)

	// History Logs
	router.GET("/api/history-logs", middlewares.AuthMiddleware(), historyControllers.ListHistoryLog)

	return router
}
