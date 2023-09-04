package main

import "time"

type UserM struct {
	ID        int64     `gorm:"column:id`
	UserID    string    `gorm:"column:user_id`
	Username  string    `gorm:"column:username`
	Status    int32     `gorm:"column:status`
	Nickname  string    `gorm:"column:nickname`
	Password  string    `gorm:"column:password`
	Email     string    `gorm:"column:email`
	Phone     string    `gorm:"column:phone`
	CreatedAt time.Time `gorm:"column:createdAt`
	UpdatedAt time.Time `gorm:"column:updatedAt`
}

func (*UserM) TableName() string {
    return "uc_user"
}

