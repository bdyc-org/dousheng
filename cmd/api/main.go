package main

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/bdyc-org/dousheng/cmd/api/handlers"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRpc()
}

func main() {
	Init()
	r := gin.New()

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.LoginParm
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			userId, _, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{Username : loginVar.UserName, Password: loginVar.PassWord})
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			return userId, nil
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	apiRouter := r.Group("/douyin")
	user1 := apiRouter.Group("/user")
	user1.GET("/", handlers.UserInfo)
	user1.POST("/login/", authMiddleware.LoginHandler)
	user1.POST("/register/", handlers.Register)
	// user1.POST("/login/", handlers.Login)

	// relation
	rela := apiRouter.Group("/relation")
	rela.Use(authMiddleware.MiddlewareFunc())
	rela.POST("/action/", handlers.Follow)
	rela.GET("/follow/list/", handlers.QueryFollow)
	rela.GET("/follower/list/", handlers.QueryFollower)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
