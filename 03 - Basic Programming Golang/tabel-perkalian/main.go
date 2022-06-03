package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d ", j*i)
		}
		fmt.Println()
	}
}

func main() {
	var number int
	fmt.Scan(&number)

	cetakTablePerkalian(number)
}
