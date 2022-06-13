package pack

import (
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

func Video(item *video.Video) *favorite.Video {
	if item == nil {
		return nil
	}
	ret := &favorite.Video{
		Author: &favorite.User{
			Id:            item.Author.Id,
			Name:          item.Author.Name,
			FollowCount:   item.Author.FollowCount,
			FollowerCount: item.Author.FollowerCount,
			IsFollow:      item.Author.IsFollow,
		},
		Id:            item.Id,
		PlayUrl:       item.PlayUrl,
		CoverUrl:      item.CoverUrl,
		FavoriteCount: item.FavoriteCount,
		CommentCount:  item.CommentCount,
		IsFavorite:    item.IsFavorite,
		Title:         item.Title,
	}
	return ret
}

func VideoList(items []*video.Video) []*favorite.Video {
	videoList := make([]*favorite.Video, 0)
	for _, item := range items {
		video := Video(item)
		videoList = append(videoList, video)
	}
	return videoList
}
