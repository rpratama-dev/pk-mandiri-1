package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if (number <= 1) {
		return false
	}

	// Iterate from 2 to the root number
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var inpNumber int
	fmt.Print("Input a number: ")
	fmt.Scanln(&inpNumber)

	if (isPrime(inpNumber)) {
		fmt.Println(inpNumber, "is primes.")
	} else {
		fmt.Println(inpNumber, "not a prime number.")
	}
}
