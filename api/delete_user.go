package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	// TODO: Xóa user khỏi DB theo ID
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"id":      id,
	})
}
