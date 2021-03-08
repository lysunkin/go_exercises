package reverse

func Reverse(input string) string {
	l := len(input)
	runes := make([]rune, l)
	for _, rune := range input {
		l--
		runes[l] = rune
	}
	return string(runes[l:])
}