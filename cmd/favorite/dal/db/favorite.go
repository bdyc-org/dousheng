package db

import (
	"context"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId  uint `gorm:"index"`
	VideoId uint
}

//favorites表的表名
func (f *Favorite) TableName() string {
	return constants.FavoriteTableName
}

func CreateFavorite(ctx context.Context, favorites []*Favorite) error {
	return MyDB.WithContext(ctx).Create(favorites).Error
}

func DeleteFavorite(ctx context.Context, userID int64, videoID int64) error {
	return MyDB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&Favorite{}).Error
}

func MGetFavorite(ctx context.Context, userID int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := MyDB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FavoriteJudge(ctx context.Context, userID int64, videoIDs []int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := MyDB.WithContext(ctx).Where("user_id = ? AND video_id in ?", userID, videoIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
