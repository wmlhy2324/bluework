package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"testlinux/gin_new/bluework/funcpackage"
	"testlinux/gin_new/bluework/middleware"
	"testlinux/gin_new/bluework/user"
	"testlinux/gin_new/bluework/utils"
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
	funcpackage.Cacheuser(Username, Pwd)

}
func login(context *gin.Context) {
	if err := context.ShouldBind(&user.MyClaims{}); err != nil {

		utils.RespFail(context, "verification failed")
		return
	}
	Username := context.PostForm("Username")
	Pwd := context.PostForm("Pwd")
	flagredis := funcpackage.Queryredis(Username)
	{
		//优先在redis里面找
		if flagredis == false {
			utils.RespFail(context, "Redis no date")
			//没有则在MySQL里面找
			flag := funcpackage.Query(Username)
			if flag == false {
				//都没有则不存在
				utils.RespFail(context, "user doesn't exists")
				return
			}
		} else {
			utils.RespSuccess(context, "Got it off redis")
		}

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

// 一个获取用户信息的接口
func getuserinfo(c *gin.Context) {
	Username := c.Query("username")
	pwd := funcpackage.Chectpwd(Username)
	flag := funcpackage.Userinfo(Username)
	if flag == false {
		utils.RespFail(c, "Redis no date")
		//不成功则去MySQL里面查找
		flag1 := funcpackage.Query(Username)

		//MySQL里面没有则返回不存在信息
		if flag1 == false {
			utils.RespFail(c, "users don't exist")
		} else {

			//重新写回redis
			funcpackage.Cacheuser(Username, pwd)
			utils.RespSuccess(c, "write redis successful")

		}
	} else {
		utils.RespSuccess(c, pwd)
	}

}

//一个点赞的接口
func like(context *gin.Context) {
	Username := context.PostForm("Username")
	flag := funcpackage.Likearticle(Username)
	if flag == false {
		utils.RespFail(context, "like fail")
	}
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

//修改密码
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
