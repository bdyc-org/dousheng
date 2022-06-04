package main

import (
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
	r := gin.Default()

	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	user1 := apiRouter.Group("/user")
	user1.GET("/", handlers.UserInfo)
	user1.POST("/register/", handlers.Register)
	user1.POST("/login/", handlers.Login)

	favorite1 := apiRouter.Group("/favorite")
	favorite1.POST("/action/", handlers.Favorite)
	favorite1.GET("/list/", handlers.FacoriteList)

	if err := r.Run(); err != nil {
		klog.Fatal(err)
	}
}
