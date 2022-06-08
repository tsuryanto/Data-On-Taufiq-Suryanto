package main

import (
	prim "dataon06/bilangan_prima/prime"
	"fmt"
)

func primeX(number int) int {
	var i = 1
	if number == 1 {
		return i + 1
	}
	for number > 1 {
		i += 2
		if i%5 != 0 || i == 5 {
			if prim.PrimeNumber(i) {
				// fmt.Println(i)
				number--
			}
		}
	}
	return i
}

func main() {
	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29
}
