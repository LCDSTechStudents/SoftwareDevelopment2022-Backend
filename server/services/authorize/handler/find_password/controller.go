package find_password

import (
	"SoftwareDevelopment-Backend/server/content"
	"github.com/gin-gonic/gin"
)

func FindPassword(ctn *content.Content) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
