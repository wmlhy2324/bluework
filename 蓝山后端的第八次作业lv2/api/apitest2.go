package api

import (
	"github.com/gin-gonic/gin"
	"testlinux/gin_new/bluework/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", Register)
	r.POST("/login", login)
	//r.POST("/modify", Editpwd)
	r.POST("/update", Editpwd)
	r.GET("/getuser", getuserinfo)
	r.POST("/like", like)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8080") // 跑在 8088 端口上
}
