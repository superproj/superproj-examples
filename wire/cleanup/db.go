package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, func(), error) {
	dsn := "zero:zero(#)666@tcp(127.0.0.1:3306)/zero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {
		sqlDB.Close()
	}

	return db, cleanFunc, nil
}
