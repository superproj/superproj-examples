package main

import "fmt"

// 定义一个结构体表示矩形
type Rectangle struct {
	width, height float64
}

// 为矩形结构体定义一个方法，用于计算面积
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// 创建一个矩形对象
	rect := Rectangle{width: 10, height: 5}

	// 调用矩形对象的方法计算面积
	area := rect.Area()

	// 输出结果
	fmt.Println("Rectangle Area:", area)
}
