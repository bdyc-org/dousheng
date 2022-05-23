package handlers

import (
	"context"
	"errors"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	errTokenExpired     error = errors.New("Token已过期,请重新登录")
	errTokenNotValidYet error = errors.New("Token无效,请重新登录")
	errTokenMalformed   error = errors.New("Token不正确,请重新登录")
	errTokenInvalid     error = errors.New("这不是一个token,请重新登录")
)

var secretKey = []byte("secret_key")

//生成token
func GenerateToken(username string) (string, error) {
	var claims MyClaims
	claims.Username = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

//解析token
func ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errTokenNotValidYet
			} else {
				return nil, errTokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errTokenInvalid
	}

	return nil, errTokenInvalid
}

//token鉴权，返回登录用户的user_id
func Authentication(c *gin.Context, claims MyClaims) (user_id int64, err error) {
	user_id, err = rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: claims.Username,
	})
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
