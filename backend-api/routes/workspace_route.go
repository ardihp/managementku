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
		route.POST("/create", ctrl.CreateWorkspace)

		detail := route.Group("/:id")
		{
			detail.GET("/", ctrl.DetailWorkspace)
			detail.POST("/assign", ctrl.AssignMember)
			detail.PUT("/", ctrl.UpdateWorkspace)
			detail.DELETE("/", ctrl.DeleteWorkspace)

			card := route.Group("/card")
			{
				card.GET("/", workspaceCardControllers.ListWorkspaceCard)
				card.POST("/create", workspaceCardControllers.CreateWorkspaceCard)
			}
		}

	}
}
