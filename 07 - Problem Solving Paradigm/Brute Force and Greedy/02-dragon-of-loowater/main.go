package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	var sum, chopped int

	for i := 0; i < len(dragonHead); i++ {
		for j := 0; j < len(knightHeight); j++ {
			if dragonHead[i] <= knightHeight[j] {
				sum += knightHeight[j]
				knightHeight = append(knightHeight[:j], knightHeight[j+1:]...)
				chopped++
				break
			}
		}
	}
	if chopped == len(dragonHead) {
		fmt.Println(sum)
	} else {
		fmt.Println("knight fail")
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fail
	DragonOfLoowater([]int{7, 3}, []int{4, 3, 1, 2}) // knight fail
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
