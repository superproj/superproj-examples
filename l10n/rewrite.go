package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.German)
	fmt.Println("Hello world")
	fmt.Printf("Hello world!")
	p.Printf("Hello world!\n")
}
