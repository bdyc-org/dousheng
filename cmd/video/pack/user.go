package pack

import (
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

func User(u *user.User) *video.User {
	if u == nil {
		return nil
	}
	ret := &video.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
	return ret
}

func UserList(us []*user.User) []*video.User {
	userList := make([]*video.User, 0)
	for _, v := range us {
		if n := User(v); n != nil {
			userList = append(userList, n)
		}
	}
	return userList
}
