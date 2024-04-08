package main

import "regexp"

var specialCharRegexp = regexp.MustCompile("[^.0-9]")

type specialChar struct {
	position position
}

func getSpecialChars(line string, curLine int) []specialChar {
	specialChars := make([]specialChar, 0)

	indices := specialCharRegexp.FindAllStringIndex(line, -1)
	for _, index := range indices {
		specialChar := specialChar{
			position: position{x: index[0], y: curLine},
		}
		specialChars = append(specialChars, specialChar)
	}

	return specialChars
}
