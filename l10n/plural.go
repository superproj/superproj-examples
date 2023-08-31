package main

import (
	"fmt"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	message.Set(language.English, "我有 %d 个苹果",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))
	message.Set(language.English, "还剩余 %d 天",
		plural.Selectf(1, "%d",
			"one", "One day left",
			"other", "%[1]d days left",
		))

}

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("我有 %d 个苹果", 1)
	fmt.Println()
	p.Printf("我有 %d 个苹果", 2)
	fmt.Println()
	p.Printf("我有 %d 个苹果", 5)
	fmt.Println()
	p.Printf("还剩余 %d 天", 1)
	fmt.Println()
	p.Printf("还剩余 %d 天", 10)
	fmt.Println()
}
