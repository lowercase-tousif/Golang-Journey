package main

import "fmt"

// We will be modifying this later
func VariableDeclaration() string {
	firstName := "Tousif"
	return firstName
}

func StringConcatenation(firstArg string, secondArg string) string {
	return firstArg + secondArg
}

func IncrementInt(num int) int {
	num = num + 1
	return num
}

// Where we will run our code
func main() {
	var age int = 27
	var name string = "Tousif"
	var hobby, job string = "coding", "student"

	fmt.Printf("Hi name is %s.\nI am %d years old.\nMy hobby is %s.\nCurrently i am %s.", name, age, hobby, job)

	var message string
	var count int
	var isActive bool

	fmt.Println("Messages:", message)
	fmt.Println("Count:", count)
	fmt.Println("isActive:", isActive)

	// testVarDec := VariableDeclaration()
	// testStringCon := StringConcatenation("Tousif", "Tasrik")
	// testIncremInt := IncrementInt(age)

	// fmt.Println("Variable Declared:",testVarDec)
	// fmt.Println("String Con:",testStringCon)
	// fmt.Println("Increment Value:",testIncremInt)

}
