package main

import (
	"fmt"
)

func main() {
	var inpNumber int
	fmt.Print("Input a number: ")
	fmt.Scanln(&inpNumber)

	if (inpNumber % 7 == 0) {
		fmt.Println(inpNumber, "is multiple 7.")
	} else {
		fmt.Println(inpNumber, "not a multiple of 7.")
	}
}
