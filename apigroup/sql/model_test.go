package sql

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestModel(t *testing.T) {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		t.Fatal(err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}
}
