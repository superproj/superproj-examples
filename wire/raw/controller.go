package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Define controller layer
type UserCenterService struct {
	biz BizFactory
}

func NewUserCenterService(db *gorm.DB) *UserCenterService {
	// 创建 Biz 层实例
	biz := NewBiz(db)

	return &UserCenterService{biz: biz}
}

// Define List User API
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
	return s.biz.Users().List(ctx, req)
}
