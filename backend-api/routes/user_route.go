package routes

import (
	userControllers "kanban-board/backend-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func userRoute(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.GET("/", userControllers.ListUser)
		user.PUT("/:id", userControllers.UpdateUser)
		user.DELETE("/:id", userControllers.DeleteUser)
	}
}
