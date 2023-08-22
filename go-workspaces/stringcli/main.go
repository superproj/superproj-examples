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
