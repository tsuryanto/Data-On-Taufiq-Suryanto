package main

import "fmt"

func MaxSequence(arr []int) int {
	var max int = arr[0]
	var a int = 0
	var sum int
	var isFirst bool

loop:
	for i := a; i < len(arr); i++ {
		if isFirst {
			sum = arr[i]
			isFirst = false
		} else {
			sum += arr[i]
		}

		if max < sum {
			max = sum
		}

		if i == len(arr)-1 {
			a++
			isFirst = true
			goto loop
		}

		if a == len(arr) {
			break
		}
	}
	return max
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   //7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   //7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   //8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    //12
}
