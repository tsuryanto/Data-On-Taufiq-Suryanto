package calculator

func Add(num ...float64) float64 {
	var result float64 = num[0]
	for i := 1; i < len(num); i++ {
		result += num[i]
	}
	return result
}

func Substract(num ...float64) float64 {
	var result float64 = num[0]
	for i := 1; i < len(num); i++ {
		result -= num[i]
	}
	return result
}

func Div(a float64, b float64) float64 {
	return a / b
}

func Multiply(num ...float64) float64 {
	var result float64 = num[0]
	for i := 1; i < len(num); i++ {
		result *= num[i]
	}
	return result
}
