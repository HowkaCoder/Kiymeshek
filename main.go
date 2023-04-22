package main

import (
	"go/todo/controllers"
	"go/todo/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Request.Header.Set("Content-Type", "application/json")
	})
	r.GET("/automigrate", controllers.Migrate)
	r.GET("/posts", controllers.PostIndex)
	r.POST("/post", controllers.PostCreate)
	r.GET("/post/:id", controllers.PostShow)
	r.GET("/armagedon", controllers.Armagedon)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
