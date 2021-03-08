// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)
// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	candidate := strings.ReplaceAll(s, "\t", " ")
	candidate = strings.ReplaceAll(candidate, "\n", " ")
	candidate = strings.ReplaceAll(candidate, "\r", " ")
	candidate = strings.ReplaceAll(candidate, ",", " ")
	candidate = strings.ReplaceAll(candidate, ".", " ")
	candidate = strings.ReplaceAll(candidate, ";", " ")
	candidate = strings.ReplaceAll(candidate, "-", " ")
	candidate = strings.ReplaceAll(candidate, "_", " ")
	candidate = strings.ReplaceAll(candidate, "!", " ")
	candidate = strings.ReplaceAll(candidate, ":", " ")
	candidate = strings.ReplaceAll(candidate, "(", " ")
	candidate = strings.ReplaceAll(candidate, ")", " ")

	parts := strings.Split(candidate, " ")

	result := make([]string, 0)

	for i := 0; i < len(parts); i++ {
		if parts[i] != "" {
			word := parts[i]
			first := word[0:1]
			result = append(result, strings.ToUpper(first))
		}
	}

	return strings.Join(result, "")
}
