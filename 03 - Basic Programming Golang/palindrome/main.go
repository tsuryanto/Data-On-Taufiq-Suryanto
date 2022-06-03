package main

import (
	"bufio"
	"fmt"
	"os"
)

func palindrome(input string) bool {
	var str []byte
	for i := len(input) - 1; i >= 0; i-- {
		str = append(str, input[i])
	}
	return string(str) == input
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kata := scanner.Text()

	if palindrome(kata) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
