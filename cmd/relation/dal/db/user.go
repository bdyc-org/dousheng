package db

import (
	"context"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

//数据库对应得结构体
type User struct {
	gorm.Model
	Id          	int64	`json:"id"`
	Name      		string	`json:"name"`
	FollowCount   	int64	`json:"follow_count"`
	FollowerCount 	int64	`json:"follower_count"`
	IsFollow		bool	`json:"is_follow"`
}

//user表的表名
func (u *User) TableName() string {
	return constants.UserTableName
}

func CreateUser(ctx context.Context, users []*User) error {
	return MyDB.WithContext(ctx).Create(users).Error
}

// 获取userList
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	var res []*User

	if err := MyDB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}