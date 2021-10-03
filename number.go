package go_helper

func NumberEvenFromInt(number int) int {
	var result int
	for value := 1; number > 0; number /= 10 {
		if digit := number % 10; digit&1 == 0 {
			result += value * digit
			value *= 10
		}
	}
	if result == 0 {
		return 100
	}
	return result
}

func NumberFirstInt(number int) int {
	if number < 10 {
		return number
	} else {
		return NumberFirstInt(number / 10)
	}
}
