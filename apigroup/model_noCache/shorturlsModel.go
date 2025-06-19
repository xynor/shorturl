package model_noCache

import (
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
func NewShorturlsModel(conn *gorm.DB) ShorturlsModel {
	return &customShorturlsModel{
		defaultShorturlsModel: newShorturlsModel(conn),
	}
}
