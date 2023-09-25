package main

import (
	"context"

	"github.com/jinzhu/copier"
)

type biz struct {
	ds *datastore
}

// 创建 Biz 层实例
func NewBiz(ds *datastore) *biz {
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
