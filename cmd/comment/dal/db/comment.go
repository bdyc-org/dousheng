package db

import (
	"context"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User_id     int64
	Video_id    int64
	Content     string `json:"content"`
	Create_date string `json:"create_date"`
}

//comment表的表名
func (c *Comment) TableName() string {
	return constants.CommentTableName
}

//发表评论
func CreatComment(ctx context.Context, c *Comment) error {
	return MyDB.WithContext(ctx).Create(c).Error
}

//删除评论
func DeleteComment(ctx context.Context, c *Comment) error {
	return MyDB.WithContext(ctx).Where("user_id = ? and video_id = ? and content = ?", c.User_id, c.Video_id, c.Content).Delete(&Comment{}).Error
}

//查看该视频的所有评论
func QueryComment(ctx context.Context, videoId int64) ([]*Comment, error) {
	var res []*Comment
	if err := MyDB.WithContext(ctx).Where("video_id = ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
