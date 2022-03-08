package main

import (
	"fmt"
	db "go-todoApp/db"
	env "go-todoApp/env"

	controllers "go-todoApp/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("hello")
	env.InitEnv()

	db.ConnectToDB()

	r:= gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	v1:= r.Group("/api")
	{
		users:= v1.Group("/user")
		{
			users.POST("/", controllers.CreateUser)
		}
	}
	r.Run()
}