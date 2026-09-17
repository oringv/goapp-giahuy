package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api/v1/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "List all users (V1)"})

	})

	r.Run(":8080")
}
