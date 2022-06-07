package pack

import (
	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

func MGetUserReq(userId int64, userIds []int64) *user.MGetUserRequest {
	return &user.MGetUserRequest{
		UserId: userId,
		UserIds: userIds,
	}
}

func User(v *user.User) *relation.User {
	return &relation.User{
		Id: v.Id,
		Name: v.Name,
		FollowCount: v.FollowCount,
		FollowerCount: v.FollowerCount,
		IsFollow: v.IsFollow,
	}
}

func UserList(us []*user.User) []*relation.User {
	userList := make([]*relation.User, 0)
	for _, v := range us {
		if n := User(v); n != nil {
			userList = append(userList, n)
		}
	}
	return userList
}

func Rela(m *db.Relation) *relation.Rela {
	return &relation.Rela{
		FollowId: 	m.Follow_id,
		FollowerId: m.Follower_id,
	}
}

func Relas(ms []*db.Relation) []*relation.Rela {
	rales := make([]*relation.Rela, 0)
	for _, m := range ms {
		if n := Rela(m); n != nil {
			rales = append(rales, n)
		}
	}
	return rales
}