package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetListUsersHandler - GET /api/v1/users
func GetListUsersHandler(c *gin.Context) {
	users := []gin.H{
		{
			"id":       1,
			"username": "user1",
			"email":    "user1@example.com",
		},
		{
			"id":       2,
			"username": "user2",
			"email":    "user2@example.com",
		},
	}

	c.JSON(http.StatusOK, users)
}

// GetUserDetailHandler - GET /api/v1/users/:id
func GetUserDetailHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"username": "user1",
		"email":    "user1@example.com",
	})
}
