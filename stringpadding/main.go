package stringpadding

import (
	"strings"
	"unicode"
)

func PadNumbers(input string, x int) string {

	if input == "" || x <= 0 {
		return input
	}

	chars := []rune(input)
	result := strings.Builder{}
	result.Grow(len(input) * 2) // preallocate

	i := 0
	for i < len(chars) {
		c := chars[i]

		if unicode.IsDigit(c) {
			// collect consecutive digits
			start := i
			for i < len(chars) && unicode.IsDigit(chars[i]) {
				i++
			}

			isFloat := false
			if i < len(chars) && hasDecimal(i, chars) {

				isFloat = true
				i++ // skip the '.'
				for i < len(chars) && unicode.IsDigit(chars[i]) {
					i++
				}
			}

			numStr := string(chars[start:i])

			if isFloat {

				// split by '.'
				parts := strings.Split(numStr, ".")
				padded := addPadding(parts[0], x)
				result.WriteString(padded)
				result.WriteRune('.')
				result.WriteString(parts[1])

			} else {
				result.WriteString(addPadding(numStr, x))
			}

		} else {
			// add non digits to the result
			result.WriteRune(c)
			i++
		}

	}

	return result.String()
}

func hasDecimal(i int, chars []rune) bool {
	return chars[i] == '.' && i+1 < len(chars) && unicode.IsDigit(chars[i+1])
}

// addPadding pads a string with leading zeros to reach length x
func addPadding(s string, x int) string {
	if len(s) >= x {
		return s
	}
	padding := strings.Repeat("0", x-len(s))
	return padding + s
}
