package main

import (
	"fmt"
)

func calculate(base, height, side float64) float64 {
	area := (base + side) * height / 2
	return area
}

func main() {
	var base, height, side float64
	fmt.Print("Enter the length of the base: ")
	fmt.Scanln(&base)
	fmt.Print("Enter the height: ")
	fmt.Scanln(&height)
	fmt.Print("Enter the length of the side: ")
	fmt.Scanln(&side)

	area := calculate(base, height, side)
	fmt.Println("The area of the trapezium is:", area)
}
