package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	// 定义一个逻辑表达式
	expression, err := govaluate.NewEvaluableExpression("a > 5 && b < 10")

	if err != nil {
		fmt.Println("表达式解析错误:", err)
		return
	}

	// 准备变量的值
	parameters := make(map[string]interface{}, 2)
	parameters["a"] = 7
	parameters["b"] = 8

	// 计算逻辑表达式
	result, err := expression.Evaluate(parameters)

	if err != nil {
		fmt.Println("表达式计算错误:", err)
		return
	}

	fmt.Println("逻辑表达式的计算结果:", result)
}
