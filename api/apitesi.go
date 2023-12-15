package api

import (
	"gin_new/bluework/funcpackage"
	"gin_new/bluework/middleware"
	"gin_new/bluework/user"
	"gin_new/bluework/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Register(context *gin.Context) {

	Username := context.PostForm("Username")
	Pwd := context.PostForm("Pwd")
	flag := funcpackage.Query(Username)
	if flag == true {
		context.JSON(500, gin.H{
			"state":   500,
			"message": "user has exists",
		})
		return
	}
	funcpackage.Addusers(Username, Pwd)
	context.JSON(200, gin.H{
		"state":   200,
		"message": "adduser successful",
	})

}
func login(context *gin.Context) {
	if err := context.ShouldBind(&user.MyClaims{}); err != nil {

		utils.RespFail(context, "verification failed")
		return
	}
	Username := context.PostForm("Username")
	Pwd := context.PostForm("Pwd")
	flag := funcpackage.Query(Username)
	if flag == false {
		utils.RespFail(context, "user doesn't exists")
		return
	}
	rightpwd := funcpackage.Checkpwd(Username)
	if rightpwd != Pwd {
		utils.RespFail(context, "worry pwd")
		return
	}
	claim := user.MyClaims{
		Username: Username,
		Pwd:      Pwd,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "李灏宇",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(context, tokenString)

}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
func Editpwd(context *gin.Context) {
	Username := context.PostForm("Username")
	Pwd := context.PostForm("Pwd")
	//想要修改的密码
	Repwd := context.PostForm("Repwd")
	flag := funcpackage.Updatepwd(Username, Pwd, Repwd)
	if flag == true {
		utils.RespSuccess(context, "Correct password")

	} else {
		utils.RespFail(context, "False password")
		return
	}
	utils.RespSuccess(context, "Update succeddfully")

}
