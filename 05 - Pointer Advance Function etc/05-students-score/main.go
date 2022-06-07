package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var sumScore int
	for _, score := range s.score {
		sumScore += score
	}
	return float64(sumScore) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	var tempScore = s.score[0]
	var tempName = s.name[0]
	for i := 0; i < len(s.score); i++ {
		if tempScore > s.score[i] {
			tempScore = s.score[i]
			tempName = s.name[i]
		}
	}
	return tempScore, tempName
}

func (s Student) Max() (max int, name string) {
	var tempScore = s.score[0]
	var tempName = s.name[0]
	for i := 0; i < len(s.score); i++ {
		if tempScore < s.score[i] {
			tempScore = s.score[i]
			tempName = s.name[i]
		}
	}
	return tempScore, tempName
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Printf("Input %d Student's Name : ", i)
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Printf("Input %s Sudent's score : ", name)
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMin, nameMin := a.Min()
	fmt.Printf("Min Score Students is : %s (%d)\n", nameMin, scoreMin)
	scoreMax, nameMax := a.Max()
	fmt.Printf("Max Score Students is : %s (%d)\n", nameMax, scoreMax)
}
