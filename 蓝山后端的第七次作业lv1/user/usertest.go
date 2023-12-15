package user

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Username string
	Pwd      string
	jwt.StandardClaims
}
