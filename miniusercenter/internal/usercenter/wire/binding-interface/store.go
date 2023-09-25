package main

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var StoreProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

type IStore interface {
	ListUser(ctx context.Context, offset, limit int) (int64, []*UserM, error)
}

type datastore struct {
	db *gorm.DB
}

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
