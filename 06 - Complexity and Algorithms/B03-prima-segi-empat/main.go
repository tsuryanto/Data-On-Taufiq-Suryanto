package main

import (
	prim "dataon06/bilangan_prima/prime"
	"fmt"
)

func primaSegiEmpat(high, wide, start int) {
	var (
		i       int = 0
		total   int
		counter int
	)
	for {
		if start >= 3 {
			start += 2
		} else {
			start += 1
		}
		if start%5 != 0 || start == 5 {
			if prim.PrimeNumber(start) {
				fmt.Printf("%d ", start)
				total += start
				i++
				counter++
			} else {
				continue
			}

			if i == high {
				i = 0
				fmt.Println()
			}
		}
		if counter == high*wide {
			fmt.Printf("%d\n\n", total)
			break
		}
	}
}

func main() {
	primaSegiEmpat(2, 3, 13)
	/*
		17 19
		23 29
		31 37
		156
	*/
	primaSegiEmpat(5, 2, 1)
	/*
		2 3 5 7 11
		13 17 19 23 19
		129
	*/
}
