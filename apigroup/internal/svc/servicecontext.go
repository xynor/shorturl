package svc

import (
	"shorturl/apigroup/internal/config"

	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.Mysql)
	if err != nil {
		logx.Must(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
