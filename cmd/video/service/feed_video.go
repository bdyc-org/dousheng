package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/cmd/video/pack"
	"github.com/bdyc-org/dousheng/cmd/video/rpc"
	favorite2 "github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type FeedVideoService struct {
	ctx context.Context
}

func NewFeedVideoService(ctx context.Context) *FeedVideoService {
	return &FeedVideoService{ctx: ctx}
}

func (v *FeedVideoService) FeedVideo(req *video.DouyinFeedRequest) ([]*video.Video, *int64, error) {
	videos, nextTime, err := db.VideoFeed(v.ctx, req.LatestTime)

	//获取author信息
	var authorsId []int64

	for i, video := range videos {
		authorsId[i] = int64(video.User_id)
	}
	resq := user.MGetUserRequest{UserId: *req.UserId, UserIds: authorsId}

	authors, err := rpc.MGetUser(v.ctx, &resq)

	var authorMap map[int64]user.User

	//将author对应的id装进map中
	for _, author := range authors {
		authorMap[author.Id] = *author
	}

	//获取点赞信息
	var videoIds []int64
	for i, video := range videos {
		videoIds[i] = int64(video.ID)
	}

	resp2 := favorite2.FavoriteJudgeRequest{UserId: *req.UserId, VideoIds: videoIds}

	favoritesvideoid, err := rpc.FavoriteJudge(v.ctx, &resp2)

	videosfinal := pack.Videos(videos)

	for _, video := range videosfinal {
		video.Author.UserId = authorMap[video.Author.UserId].Id
		video.Author.FollowerCount = authorMap[video.Author.UserId].FollowerCount
		video.Author.FollowCount = authorMap[video.Author.UserId].FollowCount
		video.Author.Name = authorMap[video.Author.UserId].Name
		video.Author.IsFollow = authorMap[video.Author.UserId].IsFollow

		video.IsFavorite = false

		for _, favoritesvideo := range favoritesvideoid {
			if video.Id == favoritesvideo {
				video.IsFavorite = true
			}
		}

		ipvf, _ := pack.GetLocalIPv4Address()

		video.PlayUrl = "http://" + ipvf + ":8080/static" + video.PlayUrl

		video.CoverUrl = "http://" + ipvf + ":8080/static" + video.CoverUrl
	}

	if err != nil {
		return nil, nil, err
	}
	return videosfinal, nextTime, nil
}
