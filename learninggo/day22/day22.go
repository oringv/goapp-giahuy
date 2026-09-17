package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, Pham Thanh Gia Huy",
		})
	})

	r.GET("/users/:user_id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Danh sach thanh vien"})
	})

	r.GET("/user/:user_id", func(ctx *gin.Context) {

		userID := ctx.Param("user_id")

		ctx.JSON(200, gin.H{
			"data":    "Thong tin user",
			"user_id": userID,
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Danh sach san pham"})
	})

	r.Run(":8080")
}
