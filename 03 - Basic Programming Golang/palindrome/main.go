package main

import (
	"fmt"
)

func palindrome(input string) bool {
	var str []byte
	for i := len(input) - 1; i >= 0; i-- {
		str = append(str, input[i])
	}
	return string(str) == input
}

func main() {
	var kata string
	fmt.Scanln(&kata)

	if palindrome(kata) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
