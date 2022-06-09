package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	for index, card := range cards {
		if deck[1] == card[0] {
			return cards[index]
		}
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino(
		[][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}},
		[]int{4, 3},
	)) // [3 4]

	fmt.Println(playingDomino(
		[][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}},
		[]int{3, 6},
	)) // [6 5]

	fmt.Println(playingDomino(
		[][]int{{6, 6}, {2, 4}, {3, 6}},
		[]int{5, 1},
	)) // "tutup kartu"
}
