package pack

import (
	"net"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

var IPAddr string

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
		IsFollow:        false,
		Avatar:          "http://" + IPAddr + ":8080/static" + u.Avatar,
		Signature:       u.Signature,
		BackgroundImage: "http://" + IPAddr + ":8080/static" + u.BackgroundImage,
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

func GetLocalIPv4Address() (err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						IPAddr = ipnet.IP.String()
					}
				}
			}
		}
	}
	return nil
}
