package pack

import (
	"time"

	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *comment.BaseResponse {
	return &comment.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}

func MGetUserReq(userId int64, userIds []int64) *user.MGetUserRequest {
	return &user.MGetUserRequest{
		UserId: userId,
		UserIds: userIds,
	}
}

func User(v *user.User) *comment.User {
	return &comment.User{
		Id: v.Id,
		Name: v.Name,
		FollowCount: v.FollowCount,
		FollowerCount: v.FollowerCount,
		IfFollow: v.IsFollow,
	}
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

func Comment(m *db.Comment, user *comment.User) *comment.Comment {
	return &comment.Comment{
		Id: m.Video_id,
		User: user,
		Content: m.Content,
		CreateDate: m.Create_date,
	}
}

func Comments(ms []*db.Comment, users []*comment.User) []*comment.Comment {
	comments := make([]*comment.Comment, 0)
	for i, m := range ms {
		if n := Comment(m, users[i]); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}