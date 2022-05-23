package db

import (
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	//连接mysql数据库
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLTestDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	//设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	//检查表是否存在，若不存在，先建表
	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}

	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}
