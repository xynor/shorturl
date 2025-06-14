package model

import (
	"fmt"
	"shorturl/rpc-gorm/transform/internal/config"
	"shorturl/rpc-gorm/transform/internal/svc"
	"testing"
)

func TestMigrate(t *testing.T) {
	sc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	err := sc.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
		AutoMigrate(&Shorturl{})
	fmt.Printf("err: %v", err)
}
