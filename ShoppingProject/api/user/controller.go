package user

import (
	"ShoppingProject/config"
	"ShoppingProject/domain/user"
	"ShoppingProject/utils/api_helper"
	jwtHelper "ShoppingProject/utils/jwt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

// 实例化
func NewUserController(service *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

// CreateUser godoc
// @Summary 根据给定的用户名和密码创建用户
// @Tags Auth
// @Accept json
// @Produce json
// @Param CreateUserRequest body CreateUserRequest true "user information"
// @Success 201 {object} CreateUserResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user [post]
func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	newUser := user.NewUser(req.Username, req.Password, req.Password2)
	err := c.userService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateUserResponse{
			Username: req.Username,
		})
}

// Login godoc
// @Summary 根据用户名和密码登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "user information"
// @Success 200 {object} LoginResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user/login [post]
func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)

	}
	currentUser, err := c.userService.GetUser(req.Username, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	decodedClaims := jwtHelper.VerifyToken(currentUser.Token, c.appConfig.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp": time.Now().Add(
					24 *
						time.Hour).Unix(),
				"isAdmin": currentUser.IsAdmin,
			})
		token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		currentUser.Token = token
		err = c.userService.UpdateUser(&currentUser)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}

	g.JSON(
		http.StatusOK, LoginResponse{Username: currentUser.Username, UserId: currentUser.ID, Token: currentUser.Token})
}

// ChangePwd godoc
// @Summary 修改密码
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization  header    string  true  "Authentication header"
// @Param ChangeRequest body ChangeRequest true "user information"
// @Success 200 {object} ChangeResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user/password [patch]
func (c *Controller) ChangePwd(g *gin.Context) {
	var req ChangeRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
	}
	userId := api_helper.GetUserId(g)
	currentUser, err := c.userService.ChangeUser(userId, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	currentUser.NewPassword = req.Password2
	err = c.userService.UpdateUser(&currentUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusOK, ChangeResponse{NewPassword: currentUser.Password})
}

// ChangeUserName godoc
// @Summary 修改用户名
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization  header    string  true  "Authentication header"
// @Param ChangeNameRequest body ChangeNameRequest true "user information"
// @Success 200 {object} ChangeNameResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user/username [patch]
func (c *Controller) ChangeUserName(g *gin.Context) {
	var req ChangeNameRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
	}

	userId := api_helper.GetUserId(g)
	currentUser, err := c.userService.ChangeName(userId, req.Username)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err = c.userService.UpdateUser(&currentUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusOK, ChangeNameResponse{NewUserName: currentUser.Username})

}

// 验证token
func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)

	g.JSON(http.StatusOK, decodedClaims)

}
