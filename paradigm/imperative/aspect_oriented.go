package main

import "fmt"

func main() {
	num1 := 5
	num2 := 3
	result := add(num1, num2)
	fmt.Println("两数之和为：", result)
}

func add(a, b int) int {
	return a + b
}
