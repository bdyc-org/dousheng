package db

import (
	"context"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	Title          string `json:"title"`
	Play_url       string `json:"play_url"`
	Cover_url      string `json:"cover_url"`
	User_id        uint   `json:"user_id"`
	Favorite_count int    `json:"favorite_count"`
	Comment_count  int    `json:"comment_count"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

// PublishVideo create video info
func PublishVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		return err
	}
	return nil
}

// QueryVideo query video by id, Limit 1
func QueryVideo(ctx context.Context, videoID uint) (*Video, error) {
	var res *Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoID)

	if err := conn.Limit(1).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

//QueryVideoList By latest_time
func VideoFeed(ctx context.Context, latest_time time.Time) (*[]Video, error) {
	var videoList *[]Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("created_at > ?", latest_time)
	if err := conn.Limit(30).Find(&videoList).Error; err != nil {
		return videoList, err
	}
	return videoList, nil
}

// DeleteVideo delete video by id
func DeleteVideo(ctx context.Context, videoID uint) error {
	return DB.WithContext(ctx).Where("id = ?", videoID).Delete(&Video{}).Error
}
