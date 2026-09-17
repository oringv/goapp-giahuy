package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Lưu vào Database
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": gin.H{
			"username":  req.Username,
			"email":     req.Email,
			"full_name": req.FullName,
		},
	})
}
