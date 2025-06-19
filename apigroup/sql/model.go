package sql

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:20"`
	Seq  string `gorm:"size:20;uniqueIndex"`
}
