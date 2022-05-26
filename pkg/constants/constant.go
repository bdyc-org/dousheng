package constants

const (
	SecretKey          = "secret key"
	IdentityKey        = "id"
	UserTableName      = "users"
	UserServiceName    = "user"
	MySQLDefaultDSN    = "jiuxia:!zzh020502@tcp(rm-bp15zhrxyp3qcn7bfto.mysql.rds.aliyuncs.com:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLTestDSN       = "root:joker@tcp(localhost:3306)/j1?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddress        = "127.0.0.1:2379"
	UserResolveTCPAddr = "127.0.0.1:6660"

	// relation
	RelationTableName  		= "relations"
	RelationServiceName    	= "relation"
	RelationResolveTCPAddr 	= "127.0.0.1:6661"
)
