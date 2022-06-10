package main

import "fmt"

func moneyCoins(money int) []int {
	var coins = []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	var exchange []int
	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] <= money {
			div := money / coins[i]
			for j := div; j > 0; j-- {
				exchange = append(exchange, coins[i])
			}
			money -= div * coins[i]
		}
		if money == 0 {
			break
		}
	}
	return exchange
}

func main() {
	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1 ]
	fmt.Println(moneyCoins(543))   // [500 20 20 111]
	fmt.Println(moneyCoins(7752))  // [5000 2000 500 200 50  1 1]
	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]
}
