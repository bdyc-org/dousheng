package constants

const (
	UserTableName                  = "users"
	UserServiceName                = "user"
	FavoriteTableName              = "favorites"
	FavoriteServiceName            = "favorite"
	RelationTableName              = "relations"
	RelationServiceName            = "relation"
	VideoTableName                 = "videos"
	VideoServiceName               = "video"
	CommentTableName               = "comments"
	CommentServiceName             = "comment"
	MySQLDefaultDSN                = "jiuxia:!zzh020502@tcp(rm-bp15zhrxyp3qcn7bfto.mysql.rds.aliyuncs.com:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLTestDSN                   = "root:123456@tcp(localhost:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddress                    = "127.0.0.1:2379"
	UserResolveTCPAddr             = "127.0.0.1:6660"
	RelationResolveTCPAddr         = "127.0.0.1:6661"
	FavoriteResolveTCPAddr         = "127.0.0.1:6602"
	CommentResolveTCPAddr          = "127.0.0.1:6603"
	VideoResolveTCPAddr			   = "127.0.0.1:6604"
	Title                          = "title"
	Videos                         = "videos"
	VideoID                        = "video_id"
	ApiServiceName                 = "demoapi"
	CPURateLimit           float64 = 80.0
	DefaultLimit                   = 10
	NextTime
)
