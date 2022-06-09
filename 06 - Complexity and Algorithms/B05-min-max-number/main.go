package main

import "fmt"

func FindMinandMax(arr []int) string {
	var iMin, iMax, min, max int = 0, 0, arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			iMin = i
			min = arr[i]
		}
		if max < arr[i] {
			iMax = i
			max = arr[i]
		}
	}
	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min, iMin, max, iMax)
}

func main() {
	fmt.Println(FindMinandMax([]int{5, 7, 4, -2, -1, 8}))
	// minL -2 index: 3 max: 8 index: 5
	fmt.Println(FindMinandMax([]int{2, -5, -4, 22, 7, 7}))
	// min: -5 index: 1 max: 22 index: 3
	fmt.Println(FindMinandMax([]int{4, 3, 9, 4, -21, 7}))
	// minL -21 index: 4 max: 9 index: 2
	fmt.Println(FindMinandMax([]int{-1, 5, 6, 4, 2, 18}))
	// minL -1 index: 0 max: 18 index: 5
	fmt.Println(FindMinandMax([]int{-2, 5, -7, 4, 7, -20}))
	// minL -20 index: 5 max: 7 index: 4
}
