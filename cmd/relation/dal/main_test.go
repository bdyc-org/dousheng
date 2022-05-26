package dal

import (
	"context"
	"testing"
	"time"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
)

var ctx context.Context
var cancel context.CancelFunc

func TestMain(m *testing.M) {
	ctx, cancel = context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	Init()
	m.Run()
}

// 找up
func TestQueryFollower(t *testing.T) {
	aaa := db.QueryFollower(ctx, 111)
	println(aaa)
}

// 找粉丝
func TestQueryFollow(t *testing.T) {
	aaa := db.QueryFollow(ctx, 111)
	println(aaa)
}

func TestFollow(t *testing.T) {
	r := db.Relation{
		Follow_id: 111,
		Follower_id: 333,
	}
	db.Follow(ctx, &r)
}

func TestCancelFollow(t *testing.T) {
	r := db.Relation{
		Follow_id: 111,
		Follower_id: 333,
	}
	db.CancelFollow(ctx, &r)
}

func TestQueryUserList(t *testing.T) {
	s := make([]int64, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	r, _ := db.MGetUsers(ctx, s)
	print(r)
}