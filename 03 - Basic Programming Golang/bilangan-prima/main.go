package main

import "fmt"

func primeNumber(number *int) bool {
	if *number < 2 {
		return false
	}
	for i := 2; i < *number; i++ {
		if *number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Scanln(&number)

	if primeNumber(&number) {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}
