package pack

import (
	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/constants"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	ret := &user.User{
		Id:              int64(u.ID),
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		Avatar:          "http://" + constants.ApiServerAddr + "/static" + u.Avatar,
		Signature:       u.Signature,
		BackgroundImage: "http://" + constants.ApiServerAddr + "/static" + u.BackgroundImage,
		TotalFavorited:  u.TotalFavorited,
		FavoriteCount:   u.FavoriteCount,
	}

	return ret
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
