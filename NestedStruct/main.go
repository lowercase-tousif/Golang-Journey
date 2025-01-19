package main

import (
	"fmt"
)

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel wheel
	BackWheel  wheel
}

type wheel struct {
	Radius   int
	Material string
}

func main() {
	myCar := car{
		Make:   "Germany",
		Model:  "xm",
		Height: 100,
		Width:  200,

		FrontWheel: wheel{
			Radius:   5,
			Material: "iron",
		},

		BackWheel: wheel{
			Radius:   5,
			Material: "iron",
		},
	}

	fmt.Println(myCar)
}
