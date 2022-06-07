package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var res []int
	var maa = make(map[int][]int)
	for index, pair := range arr {
		maa[target-pair] = []int{index, pair}
	}

	for index, pair := range arr {
		if _, ok := maa[pair]; ok {
			if maa[pair][1]+pair == target && maa[pair][1] != pair {
				res = append(res, index)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1,3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0,2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2,3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1,2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0,1]
}
