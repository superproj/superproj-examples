package main

import "fmt"

// 定义一个函数，接受两个整数并返回它们的和
func add(a, b int) int {
	return a + b
}

// 定义一个函数，接受一个函数和两个整数，并返回函数的执行结果
func compose(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main() {
	result := compose(add, 3, 5)
	fmt.Println("两数之和为：", result)
}
