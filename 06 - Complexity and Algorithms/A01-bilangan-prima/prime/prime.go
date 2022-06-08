package prime

// This function was used by another Modules
func PrimeNumber(number int) bool {
	if number == 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
