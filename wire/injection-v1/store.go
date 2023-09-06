package main

import (
	"context"

	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

// 创建 Store 层实例
func NewStore(db *gorm.DB) *datastore {
	return &datastore{db}
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
