package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var temp = make(map[string]interface{})
	for _, a := range arrayA {
		temp[a] = nil
	}
	for _, b := range arrayB {
		if _, ok := temp[b]; !ok {
			temp[b] = nil
			arrayA = append(arrayA, b)
		}
	}
	return arrayA
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"whoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
