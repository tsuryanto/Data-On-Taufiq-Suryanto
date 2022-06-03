package main

import "fmt"

func pangkat(base, pangkat *int) int {
	var total int = 1
	for {
		total *= *base
		*pangkat--
		if *pangkat == 0 {
			break
		}
	}
	return total
}

func main() {
	var x, n int
	fmt.Scan(&x)
	fmt.Scan(&n)

	fmt.Println(pangkat(&x, &n))
}
