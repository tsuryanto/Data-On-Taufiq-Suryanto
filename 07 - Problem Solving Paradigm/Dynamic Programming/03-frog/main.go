package main

import "fmt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Frog(jumps []int) int {
	var sum, i int
	var length = len(jumps)
	var lastIndex = length - 1
	for i < length {
		// fmt.Println("Path : ", i+1)
		if i < length-2 {
			if Abs(jumps[i+1]-jumps[lastIndex]) < Abs(jumps[i+2]-jumps[lastIndex]) {
				sum += Abs(jumps[i+1] - jumps[i])
				i++
			} else {
				sum += Abs(jumps[i+2] - jumps[i])
				i += 2
			}
		} else {
			if i == length-1 {
				break
			} else {
				sum += Abs(jumps[i+1] - jumps[i])
				i++
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
