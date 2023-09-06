package main

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

// 创建 DB 实例
func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error

		dsn := "zero:zero(#)666@tcp(127.0.0.1:3306)/zero?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})

	return db
}
