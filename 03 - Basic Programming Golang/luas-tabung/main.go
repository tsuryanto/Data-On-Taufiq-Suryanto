package main

import (
	"fmt"
)

func main() {
	var t, r, luasPermukaanTabung float64

	fmt.Scanf("%f", &t)
	fmt.Scanf("%f", &r)

	luasPermukaanTabung = 2*(3.14*r*r) + 2*(3.14*r*t)
	fmt.Printf("%.2f", luasPermukaanTabung)
}
