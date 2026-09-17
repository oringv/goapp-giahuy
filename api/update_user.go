package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")

	var req UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Cập nhật DB theo ID
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"id":      id,
		"data":    req,
	})
}
