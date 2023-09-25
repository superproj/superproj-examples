package service

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/superproj/superproj-examples/miniusercenter/pkg/api/miniusercenter/v1"
)

// CreateUser receives a CreateUserRequest and creates a new user record in the datastore.
// It returns an empty response (emptypb.Empty) and an error if there's any.
func (s *UserCenterService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.biz.Users().Create(ctx, req)
}
