package routes

import (
	"rest/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, routeName string) {
	userRoutes := router.Group(routeName)
	userRoutes.GET("", controllers.ListUsers)
	userRoutes.GET("/:id", controllers.GetUser)
	userRoutes.POST("", controllers.CreateUser)
	userRoutes.DELETE("/:id", controllers.DeleteUser)
	userRoutes.PATCH("/:id", controllers.UpdateUser)
}
