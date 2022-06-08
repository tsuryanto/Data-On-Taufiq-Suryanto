package main

import "fmt"

func Pow(x, n int) int {
	var r int = 1
	for n > 0 {
		if n%2 == 1 {
			r *= x
		}
		n /= 2
		x *= x
		// fmt.Printf("r: %d n: %d x: %d \n", r, n, x)
	}
	return r
}

func main() {
	fmt.Println(Pow(2, 3))  //8
	fmt.Println(Pow(5, 3))  //125
	fmt.Println(Pow(10, 2)) //100
	fmt.Println(Pow(2, 5))  //32
	fmt.Println(Pow(7, 3))  //343
}
