package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "success", "data": data})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"status": "error", "message": message})
}