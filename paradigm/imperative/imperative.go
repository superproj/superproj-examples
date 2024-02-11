package main

import "fmt"

// 过程1：计算两个数的和
func add(x, y int) int {
	return x + y
}

// 过程2：计算两个数的差
func subtract(x, y int) int {
	return x - y
}

func main() {
	// 过程式编程的调用
	a, b := 10, 5
	sum := add(a, b)       // 调用过程1
	diff := subtract(a, b) // 调用过程2

	// 输出结果
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
}
