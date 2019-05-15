package infrastructure

import (
	gin "github.com/gin-gonic/gin"
	"github.com/ryuseikurata/DDD/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}