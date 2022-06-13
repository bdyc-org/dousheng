package db

import (
	"context"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  int64
	VideoId int64 `gorm:"index"`
	Content string
}

//comment表的表名
func (c *Comment) TableName() string {
	return constants.CommentTableName
}

//发表评论
func CreateComment(ctx context.Context, item *Comment) error {
	return MyDB.WithContext(ctx).Create(&item).Error
}

//删除评论
func DeleteComment(ctx context.Context, commentID int64) error {
	return MyDB.WithContext(ctx).Where("id = ?", commentID).Delete(&Comment{}).Error
}

//查看该视频的所有评论
func MGetComment(ctx context.Context, videoID int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := MyDB.WithContext(ctx).Where("video_id = ?", videoID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
