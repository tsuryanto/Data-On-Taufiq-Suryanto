package main

import "fmt"

func caesar(offset int, input string) string {
	var st string
	offset = offset % 26

	for _, a := range input {
		intA := int(a)

		if 97 <= intA && intA <= 122 {
			sum := intA + offset
			if sum > 122 {
				st += string(rune(sum - 26))
			} else {
				st += string(rune(sum))
			}
		} else if 65 <= intA && intA <= 90 {
			sum := intA + offset
			if sum > 90 {
				st += string(rune(sum - 26))
			} else {
				st += string(rune(sum))
			}
		} else {
			st += string(a)
		}
	}
	return st
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(80, "Alta-Gol4ng~"))
}
