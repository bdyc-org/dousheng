package pack

import (
	"net"

	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

var IPAddr string

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

func Video(m *db.Video, author *video.User) *video.Video {
	if m == nil {
		return nil
	}
	return &video.Video{
		Id:            int64(m.ID),
		Author:        author,
		PlayUrl:       "http://" + IPAddr + ":8080/static/" + m.PlayUrl,
		CoverUrl:      "http://" + IPAddr + ":8080/static/" + m.CoverUrl,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		IsFavorite:    false,
		Title:         m.Title,
	}
}

func Videos(dbvideos []*db.Video, authors []*video.User) []*video.Video {
	videos := make([]*video.Video, 0)
	authorMap := make(map[int64]*video.User)
	for _, item := range authors {
		authorMap[item.Id] = item
	}
	for _, v := range dbvideos {
		video := Video(v, authorMap[v.AuthorId])
		videos = append(videos, video)
	}
	return videos
}
