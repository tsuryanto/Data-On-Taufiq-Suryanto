package main

func romanToInt(s string) int {
	var total, before int

	romanInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		r := romanInt[string(s[i])]
		if r >= before {
			total += r
		} else {
			total -= r
		}
		// fmt.Println(string(s[i]), " -- ", total)
		before = r
	}
	return total
}
