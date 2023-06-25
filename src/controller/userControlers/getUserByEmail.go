package userControlers

import "github.com/gin-gonic/gin"

func GetUserByEmail(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User Found",
	})
}
