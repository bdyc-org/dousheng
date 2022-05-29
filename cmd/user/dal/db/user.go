package db

import (
	"context"

	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"gorm.io/gorm"
)

//数据库对应得结构体
type User struct {
	gorm.Model
	Name            string
	Password        string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	Signature       string
	BackgroundImage string
	TotalFavorited  int64
	FavoriteCount   int64
}

//user表的表名
func (u *User) TableName() string {
	return constants.UserTableName
}

func CreateUser(ctx context.Context, users []*User) error {
	return MyDB.WithContext(ctx).Create(users).Error
}

func QueryUser(ctx context.Context, name string) ([]*User, error) {
	res := make([]*User, 0)
	if err := MyDB.WithContext(ctx).Where("name = ?", name).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := MyDB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func Follow(ctx context.Context, followID int64, followerID int64, actionType int64) error {
	tx := MyDB.WithContext(ctx)
	switch actionType {
	case 1:
		//被关注用户的粉丝列表加1
		err := tx.Model(&User{}).Where("ID = ?", followID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			return err
		}
		//关注操作用户的关注列表加1
		err = tx.Model(&User{}).Where("ID = ?", followerID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			return err
		}
	case 2:
		//被取关用户的粉丝列表减1
		err := tx.Model(&User{}).Where("ID = ?", followID).Update("follower_count", gorm.Expr("followr_count - ?", 1)).Error
		if err != nil {
			return err
		}
		//取关操作用户的关注列表减1
		err = tx.Model(&User{}).Where("ID = ?", followerID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			return err
		}
	default:
		return errno.ErrWrongOperation
	}
	return nil
}

func Favorite(ctx context.Context, userID int64, videoAuthor int64, actionType int64) error {
	tx := MyDB.WithContext(ctx)
	switch actionType {
	case 1:
		//被点赞用户的total_favorited加1
		err := tx.Model(&User{}).Where("ID = ?", videoAuthor).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error
		if err != nil {
			return err
		}
		//点赞操作用户的favorite_count加1
		err = tx.Model(&User{}).Where("ID = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			return err
		}
	case 2:
		//被点赞用户的total_favorited减1
		err := tx.Model(&User{}).Where("ID = ?", videoAuthor).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error
		if err != nil {
			return err
		}
		//点赞操作用户的favorite_count减1
		err = tx.Model(&User{}).Where("ID = ?", userID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			return err
		}
	default:
		return errno.ErrWrongOperation

	}
	return nil
}
