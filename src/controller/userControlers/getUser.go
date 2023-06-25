package userControlers

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User Found",
	})
}
