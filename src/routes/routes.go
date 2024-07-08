package routes

import (
	"reflect"
	"rest/controllers"
	"rest/db"

	"github.com/gin-gonic/gin"
)

func regCRUD(router *gin.RouterGroup, model interface{}) {
	router.GET("", func(c *gin.Context) {
		controllers.List(c, reflect.TypeOf(model))
	})

	router.GET("/:id", func(c *gin.Context) {
		controllers.Get(c, reflect.TypeOf(model))
	})

	router.POST("", func(c *gin.Context) {
		controllers.Create(c, reflect.TypeOf(model))
	})

	router.PATCH("/:id", func(c *gin.Context) {
		controllers.Update(c, reflect.TypeOf(model))
	})

	router.DELETE("/:id", func(c *gin.Context) {
		controllers.Delete(c, reflect.TypeOf(model))
	})
}

func RegisterUserRoutes(router *gin.Engine, routeName string) {
	userRoutes := router.Group(routeName)
	regCRUD(userRoutes, db.Person{})
}

func RegisterTaskRoutes(router *gin.Engine, routeName string) {
	taskRoutes := router.Group(routeName)
	regCRUD(taskRoutes, db.Task{})
	taskRoutes.POST("/assign", controllers.AssignUserToTask)
	taskRoutes.POST("/end", controllers.EndTask)
	taskRoutes.GET("/calculate", controllers.CalculatePayment)
}
