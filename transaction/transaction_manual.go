package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserM mapped from table <uc_user>
type UserM struct {
	ID        int64     `gorm:"column:id`
	UserID    string    `gorm:"column:user_id`
	Username  string    `gorm:"column:username`
	Status    int32     `gorm:"column:status`
	Nickname  string    `gorm:"column:nickname`
	Password  string    `gorm:"column:password`
	Email     string    `gorm:"column:email`
	Phone     string    `gorm:"column:phone`
	CreatedAt time.Time `gorm:"column:created_at`
	UpdatedAt time.Time `gorm:"column:updated_at`
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return "uc_user"
}

func main() {
	dsn := "zero:zero(#)666@tcp(127.0.0.1:3306)/zero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 开始事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 设置用户参数
	user := &UserM{
		Username: "John",
		Nickname: "john",
		Password: "zero(#)666",
		Email:    "colin404@foxmail.com",
		Phone:    "1812884xxxx",
	}
	// 创建用户
	if err := tx.Create(user).Error; err != nil {
		// 出现错误，回滚事务
		tx.Rollback()
		fmt.Println("Failed to create user:", err)
		return
	}

	// 更新用户
	user.Phone = "1812884yyyy"
	if err := tx.Save(user).Error; err != nil {
		// 出现错误，回滚事务
		tx.Rollback()
		fmt.Println("Failed to update user:", err)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}

	fmt.Println("User created and updated successfully")
}
