package db

import (
	"context"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

//数据库对应结构体
type Relation struct {
	gorm.Model
	Follow_id   	int64	// 关注者id/粉丝
	Follower_id  	int64	// 被关注者id
}

// 关系容器
var rela []Relation

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
func QueryFollower(ctx context.Context, userId int64) []int64 {
	MyDB.Where("follow_id = ?", userId).Find(&rela)

	// 已关注id切片
	f := make([]int64, len(rela))

	for i, v := range rela {
		f[i] = v.Follower_id
	}

	return f
}

// 查找粉丝
func QueryFollow(ctx context.Context, userId int64) []int64 {
	MyDB.Where("follower_id = ?", userId).Find(&rela)

	// 已关注id切片
	f := make([]int64, len(rela))

	for i, v := range rela {
		f[i] = v.Follow_id
	}

	return f
}