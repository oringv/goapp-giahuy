package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Kiểm tra username/password từ DB
	if req.Username == "admin" && req.Password == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   "fake-jwt-token-123456",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
