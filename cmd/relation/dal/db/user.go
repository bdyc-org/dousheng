package db

import (
	"context"

	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

//数据库对应得结构体
type User struct {
	gorm.Model
	Name          string
	Password      string
	FollowCount   int64
	FollowerCount int64
}

//user表的表名
func (u *User) TableName() string {
	return constants.UserTableName
}

func CreateUser(ctx context.Context, users []*User) error {
	return MyDB.WithContext(ctx).Create(users).Error
}

// 获取userList
func MGetUsers(ctx context.Context, userIDs []int64) ([]*relation.User, error) {
	res := make([]*relation.User, 0)

	if err := MyDB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}