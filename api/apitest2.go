package api

import (
	"gin_new/bluework/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", Register)
	r.POST("/login", login)
	//r.POST("/modify", Editpwd)
	r.POST("/update", Editpwd)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8080") // 跑在 8088 端口上
}
