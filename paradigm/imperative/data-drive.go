package main

import "fmt"

// 数据结构
type Person struct {
	Name string
	Age  int
}

// 处理数据的函数
func processData(person Person) {
	// 根据数据进行处理
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}

func main() {
	// 数据
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	// 遍历数据并调用处理函数
	for _, person := range people {
		processData(person)
	}
}
