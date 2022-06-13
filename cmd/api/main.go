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
	favorite1.GET("/list/", handlers.FavoriteList)

	//video
	apiRouter.GET("/feed", handlers.Feed)

	video1 := apiRouter.Group("/publish")
	video1.POST("/action/", handlers.PublishVideo)
	video1.GET("/list/", handlers.PublishList)

	//comment
	comment1 := apiRouter.Group("/comment")
	comment1.POST("/action", handlers.Comment)
	comment1.GET("/list/", handlers.CommentList)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
