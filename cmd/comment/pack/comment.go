package pack

import (
	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
)

func Comment(item *db.Comment, user *comment.User) *comment.Comment {
	return &comment.Comment{
		Id:         item.UserId,
		User:       user,
		Content:    item.Content,
		CreateDate: item.CreatedAt.Format("01-02"),
	}
}

func Comments(items []*db.Comment, users []*comment.User) []*comment.Comment {
	comments := make([]*comment.Comment, 0)
	for i, item := range items {
		if n := Comment(item, users[i]); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}
