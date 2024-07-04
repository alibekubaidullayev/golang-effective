package routes

import (
	"reflect"
	"rest/controllers"
	"rest/db"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, routeName string) {
	userRoutes := router.Group(routeName)

	userRoutes.GET("", func(c *gin.Context) {
		controllers.List(c, reflect.TypeOf(db.Person{}))
	})

	userRoutes.GET("/:id", func(c *gin.Context) {
		controllers.Get(c, reflect.TypeOf(db.Person{}))
	})

	userRoutes.POST("", func(c *gin.Context) {
		controllers.Create(c, reflect.TypeOf(db.Person{}))
	})

	userRoutes.PATCH("/:id", func(c *gin.Context) {
		controllers.Update(c, reflect.TypeOf(db.Person{}))
	})

	userRoutes.DELETE("/:id", func(c *gin.Context) {
		controllers.Delete(c, reflect.TypeOf(db.Person{}))
	})

}
