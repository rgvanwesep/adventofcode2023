package day1

func FindValue(input string) int {
	const (
		zero = byte('0')
		nine = byte('9')
	)
	var firstDigit, lastDigit int

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char >= zero && char <= nine {
			firstDigit = int(char - zero)
			break
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		char := input[i]
		if char >= zero && char <= nine {
			lastDigit = int(char - zero)
			break
		}
	}

	return firstDigit*10 + lastDigit
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += FindValue(input)
	}
	return sum
}
