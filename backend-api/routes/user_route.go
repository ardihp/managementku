package routes

import (
	"kanban-board/backend-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func userRoute(router *gin.RouterGroup, ctrl *user.UserController) {
	route := router.Group("/user")
	{
		route.GET("/", ctrl.ListUser)
		route.PUT("/:id", ctrl.UpdateUser)
		route.DELETE("/:id", ctrl.DeleteUser)
	}
}
