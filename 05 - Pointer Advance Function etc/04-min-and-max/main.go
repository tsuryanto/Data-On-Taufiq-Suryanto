package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	var (
		a *int = numbers[0]
		b *int = numbers[0]
	)
	for i := 0; i < len(numbers); i++ {
		if *a > *numbers[i] {
			a = numbers[i]
		}

		if *b < *numbers[i] {
			b = numbers[i]
		}
	}
	return *a, *b
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Printf("\n%d is the maximum number\n", max)
	fmt.Printf("%d is the minimum number\n", min)
}
