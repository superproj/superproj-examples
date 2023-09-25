package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/model"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/store"
	v1 "github.com/superproj/superproj-examples/miniusercenter/pkg/api/miniusercenter/v1"
)

var (
	ErrUserAlreadyExist = errors.NotFound(v1.ErrorReason_UserAlreadyExist.String(), "user already exist")
)

// UserBiz defines methods used to handle user request.
type UserBiz interface {
	Create(ctx context.Context, req *v1.CreateUserRequest) error
}

// userBiz struct implements the UserBiz interface and contains a store.IStore instance.
type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

// newUserBiz returns a new instance of userBiz.
func newUserBiz(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (b *userBiz) Create(ctx context.Context, req *v1.CreateUserRequest) error {
	userM := &model.UserM{
		Username: req.Username,
		Password: req.Password,
	}

	if err := b.ds.Users().Create(ctx, userM); err != nil {
		return ErrUserAlreadyExist
	}

	return nil
}
