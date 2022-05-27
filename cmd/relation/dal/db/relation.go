package db

import (
	"context"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

//数据库对应结构体
type Relation struct {
	gorm.Model
	Follow_id   	int64	`json:"follow_id"`
	Follower_id  	int64	`json:"follower_id"`
}

//relation表的表名
func (r *Relation) TableName() string {
	return constants.RelationTableName
}

// 关注操作
func Follow(ctx context.Context, r *Relation) error {
	return MyDB.WithContext(ctx).Create(r).Error
}

// 取关操作
func CancelFollow(ctx context.Context, r *Relation) error {
	return MyDB.WithContext(ctx).
	Where("follow_id = ? and follower_id = ?", r.Follow_id, r.Follower_id).
	Delete(r).Error
}

// 查找关注
func QueryFollower(ctx context.Context, userId int64) ([]*Relation, error) {
	var res []*Relation

	if err := MyDB.WithContext(ctx).
	Where("follow_id = ?", userId).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// 查找粉丝
func QueryFollow(ctx context.Context, userId int64) ([]*Relation, error) {
	var res []*Relation

	if err := MyDB.WithContext(ctx).
	Where("follower_id = ?", userId).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}