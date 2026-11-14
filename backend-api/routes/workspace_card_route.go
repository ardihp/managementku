package routes

import (
	workspaceCardControllers "kanban-board/backend-api/controllers/workspace/card"

	"github.com/gin-gonic/gin"
)

func workspaceCardRoute(router *gin.RouterGroup) {
	card := router.Group("/card/:id")
	{
		card.POST("/assign-member", workspaceCardControllers.AssignMember)
		card.PUT("/", workspaceCardControllers.UpdateWorkspaceCard)
		card.PUT("/order", workspaceCardControllers.UpdateWorkspaceCardOrder)
		card.PUT("/status-order", workspaceCardControllers.UpdateWorkspaceCardStatusOrder)
		card.DELETE("/", workspaceCardControllers.DeleteWorkspaceCard)
	}
}
