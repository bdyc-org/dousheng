package main

import (
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/handlers"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRpc()
}

func main() {
	Init()
	r := gin.New()

	apiRouter := r.Group("/douyin")
	user1 := apiRouter.Group("/user")
	user1.GET("/", handlers.UserInfo)
	user1.POST("/login/", handlers.Login)
	user1.POST("/register/", handlers.Register)

	// relation
	rela := apiRouter.Group("/relation")
	rela.POST("/action/", handlers.Follow)
	rela.GET("/follow/list/", handlers.QueryFollow)
	rela.GET("/follower/list/", handlers.QueryFollower)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
