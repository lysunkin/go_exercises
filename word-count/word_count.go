package wordcount

import (
	"strings"
	"regexp"
)

type Frequency map[string]int

func WordCount(text string) Frequency {
	m := make(Frequency)

	s := regexp.MustCompile("[\\:\\,\\.\\s\\t\\n\\!\\&\\@\\$\\%\\^\\&]+").Split(text, -1)

	for _, word := range s {
		if word != "" {
			w := strings.ToLower(word)
			w = strings.Trim(w, "'")
			m[w] = m[w] + 1
		}
	}

	return m
}
