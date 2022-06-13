package db

import (
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

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

	m := DB.Migrator()
	if m.HasTable(&Video{}) {
		return
	}
	if err = m.CreateTable(&Video{}); err != nil {
		panic(err)
	}
}
