package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	a = 20
	b = 30
)

func main() {
	fmt.Println("Working with the custom package")
	mathlib.Add(a, b)
}
