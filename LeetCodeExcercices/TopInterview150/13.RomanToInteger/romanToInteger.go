package romantointeger

import "fmt"

func RomanToInt(s string) int {
	romanToNum := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	i := 0
	fmt.Println("Roman number:", s)

	for i < len(s) {
		current := romanToNum[string(s[i])]
		next := 0

		if i+1 < len(s) {
			next = romanToNum[string(s[i+1])]
		}

		if current < next {
			result -= current
		} else {
			result += current
		}

		i++
	}
	return result
}
