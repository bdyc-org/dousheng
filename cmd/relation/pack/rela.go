package pack

import (
	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
)

func User(m *db.User) *relation.User {
	if m == nil {
		return nil
	}

	return &relation.User{
		Id:				int64(m.Id),
		Name: 			m.Name,
		FollowCount: 	m.FollowCount,
		FollowerCount: 	m.FollowerCount,
		IsFollow: 		m.IsFollow,
	}
}

func UserList(ms []*db.User) []*relation.User {
	users := make([]*relation.User, 0)
	for _, m := range ms {
		if n := User(m); n != nil {
			users = append(users, n)
		}
	}
	return users
}

func Rela(m *db.Relation) *relation.Rela {
	if m == nil {
		return nil
	}

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