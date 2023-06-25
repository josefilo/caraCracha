package userControlers

import "github.com/gin-gonic/gin"

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User Updated",
	})
}
