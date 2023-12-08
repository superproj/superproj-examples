package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 结构体表示 user 表的模型.
type User struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

// 指定 User 结构体映射的 MySQL 表名.
func (u *User) TableName() string {
	return "user"
}

// 命令行选项定义
var (
	host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "root", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// 解析命令行参数
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 连接数据库
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移模式，确保表存在
	db.AutoMigrate(&User{})

	// 创建记录
	user := User{Name: "Alice", Age: 30}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create user record: %v", err)
	}
	printUsers(db)

	// 根据查询条件查询记录，返回查询的第一条数据
	u := &User{}
	if err := db.Where("age > ?", 25).First(&u).Error; err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	// 更新记录
	u.Age = 60
	if err := db.Save(u).Error; err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	printUsers(db)

	// 根据指定的查询条件，删除记录
	if err := db.Where("age > ?", 25).Delete(&User{}).Error; err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	printUsers(db)
}

// 打印数据库 `user` 表中的记录
func printUsers(db *gorm.DB) {
	users := make([]*User, 0)
	var count int64
	d := db.Offset(0).
		Limit(2).
		Order("id desc").
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&count)
	if d.Error != nil {
		log.Fatalf("List users error: %v", d.Error)
	}

	log.Printf("User total count: %d", count)
	for _, user := range users {
		log.Printf("\tName: %s, Age: %d\n", user.Name, user.Age)
	}
}
