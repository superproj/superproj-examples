package main

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

var (
	storeOnce sync.Once
	ds        *datastore
)

type datastore struct {
	db *gorm.DB
}

// 创建 Store 层实例
func GetStore() *datastore {
	storeOnce.Do(func() {
		db := GetDB()
		ds = &datastore{db}
	})

	return ds
}

func (ds *datastore) ListUser(ctx context.Context, offset, limit int) (count int64, ret []*UserM, err error) {
	ans := ds.db.
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return count, ret, ans.Error
}
