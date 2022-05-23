package handlers

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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
		return nil, errno.ErrTokenInvalid
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errno.ErrTokenInvalid
	}

	return nil, errno.ErrTokenInvalid
}

//token鉴权，成功返回登录用户的user_id，失败返回0
func Authentication(c *gin.Context, claims MyClaims) (user_id int64, statusCode int64, err error) {
	user_id, statusCode, err = rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: claims.Username,
	})
	if err != nil {
		return 0, statusCode, err
	}
	return user_id, errno.SuccessCode, nil
}
