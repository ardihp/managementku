package routes

import (
	"kanban-board/backend-api/controllers/workspace"
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"

	"github.com/gin-gonic/gin"
)

func workspaceRoute(router *gin.RouterGroup, ctrl *workspace.WorkspaceController) {
	route := router.Group("/workspace")
	{
		route.GET("/", ctrl.ListWorkspace)
		route.GET("/:id", ctrl.DetailWorkspace)
		route.POST("/create", ctrl.CreateWorkspace)
		route.POST("/assign", ctrl.AssignMember)
		route.PUT("/:id", ctrl.UpdateWorkspace)
		route.DELETE("/:id", ctrl.DeleteWorkspace)

		card := route.Group("/:id/card")
		{
			card.GET("/", workspaceCardControllers.ListWorkspaceCard)
			card.POST("/create", workspaceCardControllers.CreateWorkspaceCard)
		}
	}
}
