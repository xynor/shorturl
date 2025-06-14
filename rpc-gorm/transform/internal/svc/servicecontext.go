package svc

import (
	"shorturl/pkg/gormz"
	"shorturl/rpc-gorm/transform/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: gormz.NewGormLogger(),
	})
	logx.Must(err)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
