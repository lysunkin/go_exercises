package luhn

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	l := len(input)
	runes := make([]rune, l)
	for _, rune := range input {
		l--
		runes[l] = rune
	}
	return string(runes[l:])
}

func Valid(input string) bool {
	s := strings.ReplaceAll(input, " ", "")

	l := len(s)
	if l < 2 {
		return false
	}

	s = Reverse(s)

	digits := make([]int, l)
	for i, rune := range s {
	        buf := make([]byte, 1)
        	_ = utf8.EncodeRune(buf, rune)
        	value, err := strconv.Atoi(string(buf))

		if err != nil {
			return false
		}

		if i % 2 == 1 {
			value = value * 2
			if value > 9 {
				value = value - 9
			}
			digits[i] = value
		} else {
			digits[i] = value
		}
	}

	sum := 0
	for _, i := range digits {
		sum += i
	}

	return (sum % 10 == 0)
}
