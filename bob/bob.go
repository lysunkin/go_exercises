// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import ( 
	"strings"
	"strconv"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	candidate := strings.ReplaceAll(remark, " ", "")
	candidate = strings.ReplaceAll(candidate, "\t", "")
	candidate = strings.ReplaceAll(candidate, "\n", "")
	candidate = strings.ReplaceAll(candidate, "\r", "")
	if candidate == "" {
		return "Fine. Be that way!"
	}

	upperCandidate := candidate

	for i := 0; i < 10; i++ {
		upperCandidate = strings.ReplaceAll(upperCandidate, strconv.Itoa(i), "")
	}

	upperCandidate = strings.ReplaceAll(upperCandidate, ",", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, ".", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, ";", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, "-", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, "!", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, ":", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, "(", "")
	upperCandidate = strings.ReplaceAll(upperCandidate, ")", "")

	if upperCandidate == strings.ToUpper(upperCandidate) && upperCandidate != "" {
		if strings.HasSuffix(upperCandidate, "?") && len(upperCandidate) != 1 {
			return "Calm down, I know what I'm doing!"
		}

		if len(upperCandidate) != 1 {
			return "Whoa, chill out!"
		}
	}

	if strings.HasSuffix(candidate, "?") {
		return "Sure."
	}

	return "Whatever."
}
