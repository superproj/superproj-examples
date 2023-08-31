package main

import (
	"fmt"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func init() {
	message.Set(language.English, "你迟了 %d 分钟。",
		catalog.Var("m", plural.Selectf(1, "%d",
			"one", "minute",
			"other", "minutes")),
		catalog.String("You are %[1]d ${m} late."))
}

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("你迟了 %d 分钟。", 1) // 打印： You are 1 minute late.
	fmt.Println()
	p.Printf("你迟了 %d 分钟。", 10) // 打印： You are 10 minutes late.
	fmt.Println()
}
