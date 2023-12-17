package main

import (
	"fmt"
)

func main() {
	A()
}

func A() {
	fmt.Println("Function A called.")
	fib(40)

	B()
	C()
}

func B() {
	fmt.Println("Function B called.")
	fib(40)
}

func C() {
	fmt.Println("Function C called.")
	fib(40)

	D()
}

func D() {
	fmt.Println("Function D called.")
	fib(40)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
