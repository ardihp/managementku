package routes

import (
	"kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"

	"github.com/gin-gonic/gin"
)

func workspaceRoute(router *gin.RouterGroup) {
	workspaceController := workspace.NewWorkspaceController()

	route := router.Group("/workspace")
	{
		route.GET("/", workspaceController.ListWorkspace)
		route.GET("/:id", workspaceController.DetailWorkspace)
		route.POST("/create", workspaceController.CreateWorkspace)
		route.POST("/assign", workspaceController.AssignMember)
		route.PUT("/:id", workspaceController.UpdateWorkspace)
		route.DELETE("/:id", workspaceController.DeleteWorkspace)

		card := route.Group("/:id/card")
		{
			card.GET("/", workspaceCardControllers.ListWorkspaceCard)
			card.POST("/create", workspaceCardControllers.CreateWorkspaceCard)
		}
	}
}
