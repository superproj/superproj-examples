package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type MapValue map[string]string

func (mv *MapValue) String() string {
	var pairs []string
	for k, v := range *mv {
		pairs = append(pairs, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(pairs, ",")
}

func (mv *MapValue) Set(value string) error {
	for _, part := range strings.Split(value, ",") {
		parts := strings.SplitN(part, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid map value: %s", value)
		}
		key := parts[0]
		val := parts[1]
		(*mv)[key] = val
	}

	return nil
}

func (mv *MapValue) Type() string {
	return "map"
}

func main() {
	// 创建一个map类型的命令行参数
	myMap := make(MapValue)
	pflag.Var(&myMap, "my-map", "a map value")

	// 解析命令行参数
	pflag.Parse()

	// 使用map参数
	fmt.Println("My Map:", myMap, len(myMap))
}
