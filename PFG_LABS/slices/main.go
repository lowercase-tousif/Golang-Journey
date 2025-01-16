package main

import (
	"fmt"
	"slices"
)

// We will be modifying this later
func CreateSlice() [] string{
	animalsSlice := []string{"dog", "cat", "monkey"}
	return animalsSlice
}

func ModifySlice(slice []string)[]string {
	slice[1] = "Tousif"
	return slice
}

func PopSliceValue(slice []string)[]string {
	slice = slices.Delete(slice,2,3)
	return slice
}

// Where we will run our code
func main() {
	numbersArray := [3]int{}

	numbersArray[0] = 1
	fmt.Println("This is my array",numbersArray) // array default value of int is 0

	numbersArray[1] = 2
	numbersArray[2] = 3

	fmt.Println(numbersArray)

	// Empty Slice
	digitSlice := []int{1,2,3}
	fmt.Println("This is my slice:",digitSlice)

	var realEmptySlice []int
	
	// Checking if the slices are empty
	if digitSlice == nil{
		fmt.Println("Digit Slice is empty")
	}

	if realEmptySlice == nil{
		fmt.Println("Real Empty slice is empty")
	}

	// length of the slice
	var lengthOfDigitSlice int
	lengthOfDigitSlice = len(digitSlice)

	fmt.Println("Length of the digit slice",lengthOfDigitSlice)

	myColors := []string{"red", "green", "yellow"}
	fmt.Println(myColors[0])
	//fmt.Println(myColors[4]) // runtime error: index out of range [4] with length 3

	// Slicing a slice
	myColors = append(myColors, "blue")
	fmt.Println(myColors)
	
	// looping through a slice
	
	for index, color := range myColors{
		fmt.Printf("Color: %s at position %d\n",color,index)
	}


	// popping a value from a slice

	myColors = myColors[:len(myColors)-1] // blue will be removed
	fmt.Println(myColors)

}
