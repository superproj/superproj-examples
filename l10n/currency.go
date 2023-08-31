package main

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	// 设置货币值
	amount := 1234.56

	// 创建一个message Printer对象，用于本地化货币金额
	p := message.NewPrinter(language.Chinese)

	// 格式化输出本地化后的货币值
	p.Printf("货币值(NarrowSymbol)：%s\n", currency.NarrowSymbol(currency.USD.Amount(amount)))
	p.Printf("货币值(NarrowSymbol)：%s\n", currency.NarrowSymbol(currency.CNY.Amount(amount)))
	p.Printf("货币值(Symbol)：%s\n", currency.Symbol(currency.USD.Amount(amount)))
	p.Printf("货币值(ISO)：%s\n", currency.ISO(currency.USD.Amount(amount)))
}
