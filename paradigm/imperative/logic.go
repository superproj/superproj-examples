package main

import "fmt"

func main() {
	age := 25

	// 使用 if-else 语句进行条件判断
	if age >= 18 {
		fmt.Println("你已成年")
	} else {
		fmt.Println("你未成年")
	}

	// 使用逻辑与(&&)运算符
	isAdult := age >= 18 && age <= 65
	if isAdult {
		fmt.Println("你是成年人")
	} else {
		fmt.Println("你不是成年人")
	}

	// 使用逻辑或(||)运算符
	isStudent := age < 25 || age > 65
	if isStudent {
		fmt.Println("你是学生")
	} else {
		fmt.Println("你不是学生")
	}
}
