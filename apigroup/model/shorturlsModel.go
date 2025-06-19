package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ShorturlsModel = (*customShorturlsModel)(nil)

type (
	// ShorturlsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShorturlsModel.
	ShorturlsModel interface {
		shorturlsModel
		customShorturlsLogicModel
	}

	customShorturlsModel struct {
		*defaultShorturlsModel
	}

	customShorturlsLogicModel interface {
	}
)

// NewShorturlsModel returns a model for the database table.
func NewShorturlsModel(conn *gorm.DB, c cache.CacheConf) ShorturlsModel {
	return &customShorturlsModel{
		defaultShorturlsModel: newShorturlsModel(conn, c),
	}
}

func (m *defaultShorturlsModel) customCacheKeys(data *Shorturls) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
