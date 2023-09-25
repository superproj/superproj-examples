// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.9
// source: miniusercenter/v1/miniusercenter.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserCenterCreateUser = "/miniusercenter.v1.UserCenter/CreateUser"

type UserCenterHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error)
}

func RegisterUserCenterHTTPServer(s *http.Server, srv UserCenterHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/users", _UserCenter_CreateUser0_HTTP_Handler(srv))
}

func _UserCenter_CreateUser0_HTTP_Handler(srv UserCenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCenterCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type UserCenterHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type UserCenterHTTPClientImpl struct {
	cc *http.Client
}

func NewUserCenterHTTPClient(client *http.Client) UserCenterHTTPClient {
	return &UserCenterHTTPClientImpl{client}
}

func (c *UserCenterHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCenterCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
