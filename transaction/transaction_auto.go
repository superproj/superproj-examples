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

	// 设置用户参数
	user := &UserM{
		Username: "John",
		Nickname: "john",
		Password: "zero(#)666",
		Email:    "colin404@foxmail.com",
		Phone:    "1812884xxxx",
	}

	// Transaction 中的所有数据库操作都处在一个事务中
	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(user).Error; err != nil {
			// 返回任何错误都会回滚事务
			fmt.Println("Failed to create user:", err)
			return err
		}

		// 更新用户
		user.Phone = "1812884yyyy"
		if err := tx.Save(user).Error; err != nil {
			// 返回任何错误都会回滚事务
			fmt.Println("Failed to update user:", err)
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	fmt.Println("User created and updated successfully")
}
