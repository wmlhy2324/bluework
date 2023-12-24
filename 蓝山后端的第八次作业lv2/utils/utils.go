package utils

import (
	"github.com/gin-gonic/gin"
)

func RespFail(c *gin.Context, message string) {

	c.JSON(400, gin.H{
		"status":  "fail",
		"message": message,
	})
}
func RespSuccess(c *gin.Context, tokenString string) {

	c.JSON(200, gin.H{
		"status": "success",
		"token":  tokenString,
	})
}
