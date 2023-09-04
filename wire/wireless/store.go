package main

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	S    *datastore
)

// Define Store layer
type IStore interface {
	Users() UserStore
}

type datastore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

func (ds *datastore) Users() UserStore {
	return newUserStore(ds.db)
}

// Define user store
type UserStore interface {
	List(ctx context.Context,  offset, limit int) (int64, []*UserM, error)
}

type userStore struct {
    db *gorm.DB
}

func newUserStore(db *gorm.DB) *userStore {
    return &userStore{db}
}


func (d *userStore) List(ctx context.Context, offset, limit int) (count int64, ret []*UserM, err error) {
	ans := d.db.
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return count, ret, ans.Error
}
