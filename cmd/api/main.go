package main

import (
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/handlers"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.Default()

	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	//user
	user1 := apiRouter.Group("/user")
	user1.GET("/", handlers.UserInfo)
	user1.POST("/login/", handlers.Login)
	user1.POST("/register/", handlers.Register)

	// relation
	rela := apiRouter.Group("/relation")
	rela.POST("/action/", handlers.Follow)
	rela.GET("/follow/list/", handlers.QueryFollow)
	rela.GET("/follower/list/", handlers.QueryFollower)

	//favorite
	favorite1 := apiRouter.Group("/favorite")
	favorite1.POST("/action/", handlers.Favorite)
	favorite1.GET("/list/", handlers.FacoriteList)

	//video
	apiRouter.GET("/feed", handlers.FeedVideo)

	video := apiRouter.Group("/publish")
	video.POST("/action/", handlers.PublishVideo)
	video.GET("/list/", handlers.PublishList)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
