package main

import (
	"fmt"
	"net/http"
	"context"
	"encoding/json"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	dsn := "zero:zero(#)666@tcp(127.0.0.1:3306)/zero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 创建 Store 层实例
	ds := NewStore(db)

	// 创建 Biz 层实例
	biz := NewBiz(ds)

	// 创建 Controller 层实例
	ucs := NewUserCenterService(biz)

	// 定义处理请求的处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := ucs.ListUser(context.Background(), &ListUserRequest{Offset:0, Limit:1})
		if err != nil {
			fmt.Fprintf(w, "ListUser failed: %v", err) // 向客户端发送响应
			return
		}

		data, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(data))
		return
	})

	// 启动HTTP服务并监听指定的端口
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("启动 HTTP 服务失败:", err)
	}
}
