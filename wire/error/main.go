package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	ucs, err := initApp()
	if err != nil {
		panic(err)
	}

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
