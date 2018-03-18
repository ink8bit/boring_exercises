package main

import (
	"fmt"
	"github.com/ink8bit/hello/greeting"
	"github.com/ink8bit/hello/words"
)

func main() {
	fmt.Println("Hello from main")
	greet := greeting.SayHello("gopher")
	fmt.Println(greet)
	fmt.Println(words.CountWords(greet))
}
