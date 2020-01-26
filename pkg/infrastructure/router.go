package infrastructure

import (
	gin "github.com/gin-gonic/gin"
	// gin "gopkg.in/gin-gonic/gin.v1"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := InitializeEvent()

	// userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
