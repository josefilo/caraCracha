package userControlers

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User Created",
	})
}
