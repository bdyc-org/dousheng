package db

import (
	"context"
	"time"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	return MyDB.WithContext(ctx).Create(videos).Error
}

func MGetVideos(ctx context.Context, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIDs) == 0 {
		return res, nil
	}

	if err := MyDB.WithContext(ctx).Where("id in ?", videoIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func PublishList(ctx context.Context, authorID int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := MyDB.WithContext(ctx).Where("author_id = ?", authorID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func Feed(ctx context.Context, lastTime time.Time) ([]*Video, int64, error) {
	res := make([]*Video, 0)
	var nextTime time.Time
	err := MyDB.WithContext(ctx).Limit(20).Order("created_at desc").Where("created_at < ?", lastTime).Find(&res).Error
	if err != nil {
		return nil, lastTime.Unix(), err
	}
	len := len(res)
	if len == 0 {
		nextTime = lastTime
	} else {
		nextTime = res[len-1].CreatedAt
	}
	return res, nextTime.Unix(), nil
}

// action==1 favorite+1  action==2 favorite-1
func Favorite(ctx context.Context, videoID int64, actionType int64) error {
	tx := MyDB.WithContext(ctx)
	switch actionType {
	case 1:
		err := tx.Model(&Video{}).Where("ID = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			return err
		}
	case 2:
		err := tx.Model(&Video{}).Where("ID = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			return err
		}
	default:
		return errno.ErrWrongOperation
	}
	return nil
}

func Comment(ctx context.Context, videoID int64, actionType int64) error {
	tx := MyDB.WithContext(ctx)
	switch actionType {
	case 1:
		err := tx.Model(&Video{}).Where("ID = ?", videoID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			return err
		}
	case 2:
		err := tx.Model(&Video{}).Where("ID = ?", videoID).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			return err
		}
	default:
		return errno.ErrWrongOperation
	}
	return nil
}
