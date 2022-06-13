package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/cmd/video/pack"
	"github.com/bdyc-org/dousheng/cmd/video/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type MGetVideoService struct {
	ctx context.Context
}

func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// 获取视频（列表）信息
func (s *MGetVideoService) MGetVideo(req *video.MGetVideoRequest) (videos []*video.Video, statusCode int64, err error) {
	modelVideos, err := db.MGetVideos(s.ctx, req.VideoIds)
	if err != nil {
		return nil, errno.ServiceErrCode, err
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
		return nil, statusCode, err
	}
	userList := pack.UserList(users)
	videos = pack.Videos(modelVideos, userList)
	//is_favorite调用favorite服务
	videoIdMap, statusCode, err := rpc.FavoriteJudge(context.Background(), &favorite.FavoriteJudgeRequest{
		UserId:   req.UserId,
		VideoIds: videoIds,
	})
	if err != nil {
		return nil, statusCode, err
	}
	for _, v := range videos {
		if _, ok := videoIdMap[v.Id]; ok {
			v.IsFavorite = true
		}
	}
	return videos, errno.SuccessCode, nil
}
