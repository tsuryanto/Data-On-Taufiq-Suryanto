package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var temp []int

loop:
	for _, a := range angka {
		val, _ := strconv.Atoi(string(a))
		for index, t := range temp {
			if t == val {
				temp = append(temp[:index], temp[index+1:]...)
				continue loop
			}
		}
		temp = append(temp, val)
	}
	return temp
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
