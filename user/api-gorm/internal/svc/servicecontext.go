package svc

import (
	"api-gorm/common/init_db"
	"api-gorm/internal/config"
	"api-gorm/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 连接数据库
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	// 连接数据表
	err := mysqlDb.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
