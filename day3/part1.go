package main

import (
	"slices"
)

type part1Parser struct {
	curLine      int
	partNumbers  []partNumber
	specialChars []specialChar
}

func (parser *part1Parser) Parse(line string) {
	parser.partNumbers = append(parser.partNumbers, getPartNumbers(line, parser.curLine)...)
	parser.specialChars = append(parser.specialChars, getSpecialChars(line, parser.curLine)...)

	parser.curLine++
}

func (parser *part1Parser) calculateResult() int {
	result := 0
	for _, partNumber := range parser.partNumbers {
		if isValidPartNumber(partNumber, parser.specialChars) {
			result += partNumber.number
		}
	}

	return result
}

func isValidPartNumber(partNumber partNumber, specialChars []specialChar) bool {
	return slices.ContainsFunc(specialChars, func(specialChar specialChar) bool {
		return specialChar.position.isInside(partNumber.topLeft, partNumber.bottomRight)
	})
}
