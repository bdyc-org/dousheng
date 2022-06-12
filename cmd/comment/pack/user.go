package pack

import (
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

func User(u *user.User) *comment.User {
	if u == nil {
		return nil
	}
	ret := &comment.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
	return ret
}

func UserList(us []*user.User) []*comment.User {
	userList := make([]*comment.User, 0)
	for _, v := range us {
		if n := User(v); n != nil {
			userList = append(userList, n)
		}
	}
	return userList
}
