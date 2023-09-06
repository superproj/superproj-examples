package main

import (
	"context"
	"fmt"
	"time"
)

type UserCenterService struct {
	biz *biz
}

// 创建 Controller 层实例
func NewUserCenterService(biz *biz) *UserCenterService {
	return &UserCenterService{biz: biz}
}

type ListUserRequest struct {
	Limit  int
	Offset int
}

type UserReply struct {
	UserID    string    `json:"userID"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListUserResponse struct {
	TotalCount int64        `json:"totalCount"`
	Users      []*UserReply `json:"users"`
}

func (s *UserCenterService) ListUser(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	fmt.Println("ListUser function called")
	return s.biz.ListUser(ctx, req)
}
