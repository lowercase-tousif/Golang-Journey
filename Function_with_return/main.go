package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func calc(num1 int , num2 int)(int , int){
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func main(){
	a := 10
	b := 20

	sum := add(a, b)

	fmt.Println(sum)

	p,q := calc(a,b)
	fmt.Println(p,q)
}