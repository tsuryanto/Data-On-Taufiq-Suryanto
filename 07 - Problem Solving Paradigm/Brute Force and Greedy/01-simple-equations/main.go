package main

import (
	"fmt"
	"math"
)

/*
	temp = c
	loop 3x :
		temp = temp - int(log2(temp))
	Di loop yang terakhir kalau hasilnya 1, return 1 (1^2 = 1)
	else if di cek lagi log2(temp) nya . Kalau bukan desimal return nilai integernya
	else (kalau desimal) return 0, berarti no solution
*/

func kuadrat(n int) int {
	return int(math.Pow(float64(n), 2))
}

func isTestPassed(a, b, c, x, y, z int) bool {
	var passed int
	if x+y+z == a {
		passed++
	}
	if x*y*z == b {
		passed++
	}
	if kuadrat(x)+kuadrat(y)+kuadrat(z) == c {
		passed++
	}
	return passed == 3
}

func SimpleEquations(a, b, c int) {
	if a < 3 {
		fmt.Println("no solution")
	} else {
		var res, x, y, z int
		var temp float64 = float64(c)
		// loop 3x
		for i := 2; 0 <= i; i-- {
			res = int(math.Log2(temp))
			if i == 0 {
				if temp == 1 { // karena 1^2 = 1
					res = 1
				} else if math.Log2(temp) != float64(res) { // kalau log 2 nya desimal , maka no solution
					res = 0
				}
			} else {
				temp = temp - math.Pow(float64(res), 2)
			}
			switch i {
			case 0:
				x = res
			case 1:
				y = res
			case 2:
				z = res
			}
		}

		if x > 0 && isTestPassed(a, b, c, x, y, z) {
			fmt.Printf("%d %d %d\n", x, y, z)
		} else {
			fmt.Println("no solution")
		}
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
	// SimpleEquations(6, 7, 14) // no solution
}
