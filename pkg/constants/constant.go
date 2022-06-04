package constants

const (
	UserTableName          = "users"
	UserServiceName        = "user"
	FavoriteTableName      = "favorites"
	FavoriteServiceName    = "favorite"
	MySQLDefaultDSN        = "jiuxia:!zzh020502@tcp(rm-bp15zhrxyp3qcn7bfto.mysql.rds.aliyuncs.com:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLTestDSN           = "root:123456@tcp(localhost:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddress            = "127.0.0.1:2379"
	UserResolveTCPAddr     = "127.0.0.1:6660"
	FavoriteResolveTCPAddr = "127.0.0.1:6602"
)
