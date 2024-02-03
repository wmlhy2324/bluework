package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 错误处理
func HandleError(c *gin.Context, err error) {

	c.JSON(
		http.StatusBadRequest,
		ErrorResponse{
			Message: err.Error(),
		})
	c.Abort()
	return
}
