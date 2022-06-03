package main

import "fmt"

func playWithAsterik(n int) {
	var x string
	if n%2 == 0 {
		x = " "
	} else {
		x = "*"
	}

	for i := n; i > 0; i-- {
		for j := 1; j < n*2; j++ {
			if j >= i && j <= n*2-i {
				fmt.Print(x)
			} else {
				fmt.Print(" ")
			}

			if x == "*" {
				x = " "
			} else {
				x = "*"
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	var height int
	fmt.Scan(&height)
	playWithAsterik(height)
}
