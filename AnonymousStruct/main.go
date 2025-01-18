package main

import "fmt"

func main() {
	person := struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}{
		FirstName: "Tousif",
		LastName:  "Tasrik",
		Age:       21,
		Email:     "tousif@gmail.com",
	}

	fmt.Println("Name: ", person.FirstName, person.LastName)
	fmt.Println("Age: ", person.Age)
	fmt.Println("Email: ", person.Email)

}
