package main

import "fmt"

func fibonacci(number int) int {
	if number == 0 {
		return 0
	} else {
		return fibo(number-1, 0, 1)
	}
}

func fibo(max int, m int, n int) int {
	// result = m + n
	if max > 1 {
		return fibo(max-1, n, m+n)
	} else {
		return m + n
	}
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
