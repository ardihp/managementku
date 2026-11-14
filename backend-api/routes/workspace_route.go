package routes

import (
	workspaceControllers "kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"

	"github.com/gin-gonic/gin"
)

func workspaceRoute(router *gin.RouterGroup) {
	workspace := router.Group("/workspace")
	{
		workspace.GET("/", workspaceControllers.ListWorkspace)
		// workspace.GET("/:id", workspaceControllers.ListWorkspace)
		workspace.POST("/create", workspaceControllers.CreateWorkspace)
		workspace.POST("/assign", workspaceControllers.AssignMember)
		workspace.PUT("/:id", workspaceControllers.UpdateWorkspace)
		workspace.DELETE("/:id", workspaceControllers.DeleteWorkspace)

		card := workspace.Group("/:id/card")
		{
			card.GET("/", workspaceCardControllers.ListWorkspaceCard)
			card.POST("/create", workspaceCardControllers.CreateWorkspaceCard)
		}
	}
}
