package db

import (
	"context"
	"fmt"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
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

// PublishVideo create videos info
func PublishVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		return err
	}
	return nil
}

// QueryVideo query videos by id, Limit 1
func QueryVideo(ctx context.Context, videoID uint) (*Video, error) {
	var res *Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoID)

	if err := conn.Limit(1).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

//QueryVideoList By latest_time
func VideoFeed(ctx context.Context, LatestTime *int64) ([]*Video, *int64, error) {

	var videoList []*Video //transfrom
	conn := DB.WithContext(ctx).Model(&Video{}).Where("created_at < FROM_UNIXTIME(?)", *LatestTime)

	err := conn.Limit(20).Find(videoList).Error

	fmt.Println(videoList[0].Title)

	var nextTime = videoList[0].Model.CreatedAt.Unix()

	if err != nil {
		return videoList, &nextTime, err
	}
	return videoList, &nextTime, nil
}

// DeleteVideo delete videos by id
func DeleteVideo(ctx context.Context, videoID uint) error {
	return DB.WithContext(ctx).Where("id = ?", videoID).Delete(&Video{}).Error
}

// QueryVideo query videos by user_id
func QueryVideos(ctx context.Context, user_id uint) ([]*Video, error) {
	var res []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("user_id = ?", user_id)

	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// action==1 favorite+1  action==2 favorite-1
func VideoFavorite(ctx context.Context, videoID uint, action int) error {
	if action == 1 {
		err := DB.WithContext(ctx).Where("id = ?", videoID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		err := DB.WithContext(ctx).Where("id = ?", videoID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	}
}

// action==1 comment+1  action==2 comment-1
func VideoComment(ctx context.Context, videoID uint, action int) error {
	if action == 1 {
		err := DB.WithContext(ctx).Where("id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		err := DB.WithContext(ctx).Where("id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	}
}
