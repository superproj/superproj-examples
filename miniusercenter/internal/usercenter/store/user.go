package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/model"
)

// UserStore defines the interface for managing user data storage.
type UserStore interface {
	// Create adds a new user record to the database.
	Create(ctx context.Context, user *model.UserM) error
}

// userStore is an implementation of the UserStore interface using a datastore.
type userStore struct {
	db *gorm.DB
	uc *userCache
}

// newUserStore returns a new instance of userStore with the provided datastore.
func newUserStore(db *gorm.DB) *userStore {
	once.Do(func() {
		us = &userStore{
			db: db,
			uc: newUserCache(),
		}
	})

	return us
}

// Create adds a new user record to the database.
func (d *userStore) Create(ctx context.Context, user *model.UserM) error {
	return d.uc.Save(user)
	// return d.db.Create(&user).Error
}
