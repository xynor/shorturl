package model

import "gorm.io/gorm"

type Shorturl struct {
	gorm.Model
	Shorten string `gorm:"size:255;uniqueIndex"`
	Url     string `gorm:"size:255"`
}
