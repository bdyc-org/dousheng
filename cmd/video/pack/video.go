package pack

import (
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type User struct {
	UserId        int64  `thrift:"user_id,1" json:"user_id"`
	Name          string `thrift:"name,2" json:"name"`
	FollowCount   int64  `thrift:"follow_count,3" json:"follow_count"`
	FollowerCount int64  `thrift:"follower_count,4" json:"follower_count"`
	IsFollow      bool   `thrift:"is_follow,5" json:"is_follow"`
}

func Video(m *db.Video) *video.Video {
	if m == nil {
		return nil
	}
	return &video.Video{
		Id:            int64(m.ID),
		Title:         m.Title,
		PlayUrl:       m.Play_url,
		CoverUrl:      m.Cover_url,
		Author:        nil, //需要用到user服务
		FavoriteCount: int64(m.Favorite_count),
		CommentCount:  int64(m.Comment_count),
		IsFavorite:    false, //需要用到点赞服务
	}
}

func Videos(dbvideos []*db.Video) []*video.Video {

	videos := make([]*video.Video, 0)
	for _, videoz := range dbvideos {
		if vide := Video(videoz); vide != nil {
			author := &video.User{
				UserId:        0,
				Name:          "",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			}
			vide.Author = author
			vide.Author.UserId = int64(videoz.User_id)
			videos = append(videos, vide)
		}
	}

	return videos
}
