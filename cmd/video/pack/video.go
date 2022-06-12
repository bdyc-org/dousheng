package pack

import (
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

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

	videos := make([]*video.Video, 30)
	for _, videoz := range dbvideos {
		video := Video(videoz)
		video.Author.UserId = int64(videoz.User_id)
		videos = append(videos, video)
	}
	return videos
}
