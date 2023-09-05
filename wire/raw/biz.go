package main

import (
	"context"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// Define biz layer
type BizFactory interface {
	Users() UserBiz
}

type biz struct {
	ds IStore
}

func NewBiz(db *gorm.DB) *biz {
	// 创建 Store 层实例
    ds := NewStore(db)

	return &biz{ds: ds}
}

func (b *biz) Users() UserBiz {
	return NewUser(b.ds)
}

// Define user biz
type UserBiz interface {
	List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error)
}

type userBiz struct {
	ds IStore
}

func NewUser(ds IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (b *userBiz) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	count, list, err := b.ds.Users().List(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	users := make([]*UserReply, 0)
	for _, item := range list {
		var u UserReply
		_ = copier.Copy(&u, &item)
		users = append(users, &u)
	}

	return &ListUserResponse{TotalCount: count, Users: users}, nil
}
