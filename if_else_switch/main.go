package main 

import "fmt"

func main(){
	age := 18

	if age > 18{
		fmt.Println("You can vote")
	}else if age < 18{
		fmt.Println("You can't vote")
	}else{
		fmt.Println("Okay let's make birth certificate")
	}

	// Important operators to know
	/*
		> ,>= , <, <=, ==
		and ==> &&
		or ==> ||
		not ==> ! 
	*/

	age1 := 20
	sex := "female"

	if age1 > 18 && sex == "female"{
		fmt.Println("You are ready to get married girl")
	}else{
		fmt.Println("Problemm")
	}

	isPretty := false

	if !isPretty{
		fmt.Println("Print somethinggg")
	}

	
	y := 2

	switch y{
	case 1:
		fmt.Println("This is 1")
	case 2,3:
		fmt.Println("This is 2 or 3")
	default:
		fmt.Println("y is neither 1 nor 2 or 3 ")
	}
}