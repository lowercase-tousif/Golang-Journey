package main

import "fmt"

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

func main() {
	bmw := car{
		Make:   "Germany",
		Model:  "i8",
		Height: 100,
		Width:  200,
	}

	fmt.Println(bmw)

	// accessing the structs value
	fmt.Println("Car's model: ", bmw.Model)
	fmt.Println("Car's height: ", bmw.Height)

	// modifying a struct
	bmw.Model = "xm"

	fmt.Println("New updated model: ", bmw.Model)
	
	
}
