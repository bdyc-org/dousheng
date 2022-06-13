package service

import (
	"context"
	"time"

	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/cmd/video/pack"
	"github.com/bdyc-org/dousheng/cmd/video/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *video.FeedRequest) (videos []*video.Video, nextTime int64, statusCode int64, err error) {
	modelVideos, nextTime, err := db.Feed(s.ctx, time.Unix(req.LatestTime, 0))
	if err != nil {
		return nil, req.LatestTime, errno.ServiceErrCode, err
	}
	userIds := make([]int64, 0)
	videoIds := make([]int64, 0)
	for _, item := range modelVideos {
		userIds = append(userIds, item.AuthorId)
		videoIds = append(videoIds, int64(item.ID))
	}
	users, statusCode, err := rpc.MGetUser(context.Background(), &user.MGetUserRequest{
		UserId:  req.UserId,
		UserIds: userIds,
	})
	if err != nil {
		return nil, 0, statusCode, err
	}
	userList := pack.UserList(users)
	videos = pack.Videos(modelVideos, userList)
	//is_favorite调用favorite服务
	if req.UserId != 0 {
		videoIdMap, statusCode, err := rpc.FavoriteJudge(context.Background(), &favorite.FavoriteJudgeRequest{
			UserId:   req.UserId,
			VideoIds: videoIds,
		})
		if err != nil {
			return nil, req.LatestTime, statusCode, err
		}
		for _, v := range videos {
			if _, ok := videoIdMap[v.Id]; ok {
				v.IsFavorite = true
			}
		}
	}
	return videos, nextTime, errno.SuccessCode, nil
}
