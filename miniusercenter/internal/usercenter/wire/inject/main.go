package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

	ds := NewStore(db)
	biz := NewBiz(ds)
	ucs := NewUserCenterService(biz)

	// 定义处理请求的处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := ucs.ListUser(context.Background(), &ListUserRequest{Offset: 0, Limit: 1})
		if err != nil {
			fmt.Fprintf(w, "ListUser failed: %v", err) // 向客户端发送响应
			return
		}

		data, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(data))
		return
	})

	// 启动HTTP服务并监听指定的端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("启动 HTTP 服务失败:", err)
	}
}
