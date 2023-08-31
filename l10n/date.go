package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 格式化输出当前时间
	fmt.Printf("当前时间：%s\n\n", now.Format("2006-01-02 15:04:05"))

	// 设置时区为中国标准时间
	locCN, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("设置时区失败：", err)
		return
	}

	// 在中国标准时间下输出当前时间
	nowInChina := now.In(locCN)
	fmt.Printf("当前时间（中国标准时间）：%s\n", nowInChina.Format("2006-01-02 15:04:05"))

	locUS, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("设置时区失败：", err)
		return
	}
	fmt.Printf("当前时间（美国标准时间）：%s\n\n", now.In(locUS).Format("2006-01-02 15:04:05"))

	// 解析日期字符串为时间
	dateStr := "2023-08-30 23:43:52"
	t, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		fmt.Println("解析日期字符串失败：", err)
		return
	}

	// 在中国标准时间下输出解析后的时间
	tInCN := t.In(locCN)
	fmt.Printf("解析后的时间（中国标准时间）：%s\n", tInCN.Format("2006-01-02 15:04:05"))
	tInUS := t.In(locUS)
	fmt.Printf("解析后的时间（美国标准时间）：%s\n", tInUS.Format("2006-01-02 15:04:05"))
}
