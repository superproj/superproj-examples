// Package main main 文件，go 多模块工作区演示代码
// 实现将输入的字符串反转输出并输出
package main

import (
	"flag"
	"fmt"

	stringsutil "github.com/superproj/superproj-examples/go-workspaces/pkg/util/strings"
)

var str string

func main() {
	flag.StringVar(&str, "str", str, "String used to deal.")
	flag.Parse()

	if str == "" {
		fmt.Println("Must specified a string")
		flag.Usage()
		return
	}

	fmt.Println(stringsutil.Reversal(str))
}
