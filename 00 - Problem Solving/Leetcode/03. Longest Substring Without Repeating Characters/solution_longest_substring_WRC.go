package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var a int = 0
	var longest int
	var sub string
	// var testStrLongest string

loop:
	for i := a; i < len(s); i++ {
		isContain := strings.Contains(sub, string(s[i]))

		sub = s[a : i+1]
		if isContain {
			if a < len(s)-1 {
				a++
				sub = ""
				goto loop
			}
		} else {
			sub = s[a : i+1]
			if longest < len(sub) {
				longest = len(sub)
				// testStrLongest = sub
			}
			// fmt.Printf("sub %s, longest %d, strLongest %s\n", sub, longest, testStrLongest)
		}
	}

	// fmt.Println("__", longest)
	return longest

}
