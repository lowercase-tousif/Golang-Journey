package main

import (
	"fmt"
)

var a = 20
var b = 30

func add(x int, y int){
	z := x + y
	fmt.Println(z)
}

func main() {
	
	var p int = 40
	var q int = 50

	add(p,q)
	add(a,b)
	add(b,q)
	
	//add(b,z) // z is not defined in this scope
}