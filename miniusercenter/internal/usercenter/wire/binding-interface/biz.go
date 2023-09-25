package main

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

var BizProviderSet = wire.NewSet(NewBiz, wire.Bind(new(BizFactory), new(*biz)))

type BizFactory interface {
	ListUser(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error)
}

type biz struct {
	ds IStore
}

func NewBiz(ds IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) ListUser(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	count, list, err := b.ds.ListUser(ctx, req.Offset, req.Limit)
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
