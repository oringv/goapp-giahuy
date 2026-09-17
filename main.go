package main

import (
	"goapp-giahuy/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", api.CreateUserHandler)       // Create
		v1.PUT("/users/:id", api.UpdateUserHandler)    // Update
		v1.DELETE("/users/:id", api.DeleteUserHandler) // Delete
		v1.GET("/users", api.GetListUsersHandler)      // Get List
		v1.GET("/users/:id", api.GetUserDetailHandler) // Get Detail
		v1.POST("/login", api.LoginHandler)            // Login
	}

	r.Run(":8080")
}
