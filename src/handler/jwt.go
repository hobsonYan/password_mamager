package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dianjiu/gokit/slice"
	"github.com/gin-gonic/gin"
)

type UserClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`

	jwt.StandardClaims
}

var (
	secret = []byte("16849841325189456f487")

	// 该路由下不校验token
	noVerify = []interface{}{"/gin/user/login", "/gin/user/register", "/gin/user/netLogin"}

	effectTime = 2 * time.Hour
)

func GenerateToken(claims *UserClaims) string {

	claims.ExpiresAt = time.Now().Add(effectTime).Unix()

	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)

	if err != nil {
		panic(err)
	}

	return sign
}

func JwtVerify(c *gin.Context) {

	if slice.InSliceIface(c.Request.RequestURI, noVerify) {
		return
	}
	token := c.GetHeader("token")
	if token == "" {
		panic("token not exist!")
	}
	c.Set("user", parseToken(token))
}

func parseToken(tokenString string) *UserClaims {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is invalid")
	}
	return claims
}

func RefreshToken(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is invalid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}
