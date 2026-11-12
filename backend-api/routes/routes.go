package routes

import (
	authControllers "kanban-board/backend-api/controllers/auth"
	historyControllers "kanban-board/backend-api/controllers/history"
	socketControllers "kanban-board/backend-api/controllers/socket"
	userControllers "kanban-board/backend-api/controllers/user"
	workspaceControllers "kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"
	workspaceCardActivityControllers "kanban-board/backend-api/controllers/workspace/card/activity"
	workspaceCardStatusControllers "kanban-board/backend-api/controllers/workspace/card/status"

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
	router.GET("/api/users", middlewares.AuthMiddleware(), userControllers.ListUser)
	router.PUT("/api/user/:id", middlewares.AuthMiddleware(), userControllers.UpdateUser)
	router.PUT("/api/user-delete/:id", middlewares.AuthMiddleware(), userControllers.DeleteUser)
	router.GET("/api/employees", middlewares.AuthMiddleware(), userControllers.ListEmployee)
	router.GET("/api/card-member/:id", middlewares.AuthMiddleware(), userControllers.ListWorkspaceCardMember)
	router.GET("/api/workspace-member/:id", middlewares.AuthMiddleware(), userControllers.ListWorkspaceMember)

	// Workspace
	router.GET("/api/workspace", middlewares.AuthMiddleware(), workspaceControllers.ListWorkspace)
	router.GET("/api/workspace/:id", middlewares.AuthMiddleware())
	router.POST("/api/workspace/create", middlewares.AuthMiddleware(), workspaceControllers.CreateWorkspace)
	router.PUT("/api/workspace/update/:id", middlewares.AuthMiddleware(), workspaceControllers.UpdateWorkspace)
	router.POST("/api/workspace/assign", middlewares.AuthMiddleware(), workspaceControllers.AssignMemberToWorkspace)
	router.DELETE("/api/workspace/delete/:id", middlewares.AuthMiddleware(), workspaceControllers.DeleteWorkspace)

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
	router.POST("/api/workspace/:id/card/assign", middlewares.AuthMiddleware(), workspaceCardControllers.AssignMemberToWorkspaceCard)
	router.PUT("/api/workspace/:id/card/watch/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.WatchWorkspaceCard)
	router.PUT("/api/workspace/:id/card/unwatch/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.UnwatchWorkspaceCard)
	router.PUT("/api/workspace/:id/card/complete/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.MarkWorkspaceCardComplete)
	router.PUT("/api/workspace/:id/card/incomplete/:card_id", middlewares.AuthMiddleware(), workspaceCardControllers.MarkWorkspaceCardIncomplete)

	// Workspace Card Status
	router.GET("/api/workspace/:id/status", middlewares.AuthMiddleware(), workspaceCardStatusControllers.ListWorkspaceCardStatus)
	router.GET("/api/workspace/:id/status/archived", middlewares.AuthMiddleware(), workspaceCardStatusControllers.ListArchivedWorkspaceCardStatus)
	router.POST("/api/workspace/:id/status/create", middlewares.AuthMiddleware(), workspaceCardStatusControllers.CreateWorkspaceCardStatus)
	router.PUT("/api/workspace/:id/status/update/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.UpdateWorkspaceCardStatus)
	router.PUT("/api/workspace/:id/status/archive/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.ArchieveWorkspaceCardStatus)
	router.PUT("/api/workspace/:id/status/unarchive/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.UnArchieveWorkspaceCardStatus)
	router.DELETE("/api/workspace/:id/status/delete/:status_id", middlewares.AuthMiddleware(), workspaceCardStatusControllers.DeleteWorkspaceCardStatus)

	// Workspace Card Activity
	router.GET("/api/workspace/:id/card/:card_id/activity", middlewares.AuthMiddleware(), workspaceCardActivityControllers.ListWorkspaceCardActivity)
	router.POST("/api/workspace/:id/card/:card_id/create-activity", middlewares.AuthMiddleware(), workspaceCardActivityControllers.CreateWorkspaceCardActivity)
	router.PUT("/api/workspace/:id/card/:card_id/update-activity/:activity_id", middlewares.AuthMiddleware(), workspaceCardActivityControllers.UpdateWorkspaceCardActivity)
	router.DELETE("/api/workspace/:id/card/:card_id/delete-activity/:activity_id", middlewares.AuthMiddleware(), workspaceCardActivityControllers.DeleteWorkspaceCardActivity)

	// History Logs
	router.GET("/api/history-logs", middlewares.AuthMiddleware(), historyControllers.ListHistoryLog)
	router.GET("/api/workspace/:id/history-logs", middlewares.AuthMiddleware(), historyControllers.ListWorkspaceHistoryLog)

	return router
}
