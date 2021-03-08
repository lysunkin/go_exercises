package isogram

import "unicode"

func IsIsogram(input string) bool {
	mapping := make(map[rune]int)
	for _, c := range input {
		if c != ' ' && c != '-' {
			mapping[unicode.ToUpper(c)]++
		}
	}

	for _, i := range mapping {
		if i != 1 {
			return false
		}
	}

	return true
}