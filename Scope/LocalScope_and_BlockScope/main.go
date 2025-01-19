package main

import (
	"fmt"
)

// global scope
var (
	a = 20
	b = 30
)

func main() {
	x := 10 // local scope

	if x > 5 {
		p := 10 // block scope
		fmt.Println("Hello How are you ?")
		fmt.Println("I have", p, "mangoes")
	}
}
