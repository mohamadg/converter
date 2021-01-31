// Package toInt converts Strings to numbers

package toFloat

import (
	"math"
	"strings"
)

// StrToInt takes a string, converts it to float64 if possible
// then returns the number
func toFloat(str string) float64 {
	// possible numbers
	numbers := map[string]float64{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	lenStr := len(str) - 1
	// this would change to true if the number is decimal
	decimal := false
	// this would change to true if the number is negative
	nagative := false
	// where indicate where the dot is
	dotPos := 0
	// the finall result
	var result float64 = 0
	// check if the string is empty
	if lenStr == -1 {
		panic("string is empty")
	}
	// check if the string contains any digit
	for i := range numbers {
		if strings.Contains(str, i) {
			break
		}
		if i == "9" {
			panic("string does not have any digit")
		}
	}
	// check if the string contains a decimal number
	if strings.Contains(str, ".") {
		pointsCount := strings.Count(str, ".")
		if pointsCount != 1 {
			panic("string is not valid")
		}
		lenStr -= 1 // lenStr minus 1, because one of chars in str is a dot
		decimal = true
	}
	for i, chr := range str {
		value, exist := numbers[string(chr)]
		if !exist {
			switch {
			case string(chr) == ".":
				dotPos = len(str) - i - 1
				continue
			case i == 0 && string(chr) == "-":
				nagative = true
			default:
				panic("string is not valid")
			}
		}
		result += value * (math.Pow(10, float64(lenStr)))
		lenStr -= 1
	}
	if nagative {
		result *= -1
	}
	if decimal {
		result /= math.Pow(10, float64(dotPos))
	}
	return result
}
