package main

import (
	"regexp"
	"slices"
	"util"
)

type position struct {
	x int
	y int
}

type partNumber struct {
	number      int
	topLeft     position
	bottomRight position
}

type specialChar struct {
	position position
}

type part1Parser struct {
	curLine      int
	partNumbers  []partNumber
	specialChars []specialChar
}

var numberRegexp = regexp.MustCompile("[0-9]+")
var specialCharRegexp = regexp.MustCompile("[^.0-9]")

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

func getPartNumbers(line string, curLine int) []partNumber {
	partNumbers := make([]partNumber, 0)

	numberIndices := numberRegexp.FindAllStringIndex(line, -1)
	for _, index := range numberIndices {
		partNumber := partNumber{
			number:      util.MustParseInt(line[index[0]:index[1]]),
			topLeft:     position{x: index[0] - 1, y: curLine - 1},
			bottomRight: position{x: index[1], y: curLine + 1},
		}
		partNumbers = append(partNumbers, partNumber)
	}

	return partNumbers
}

func isValidPartNumber(partNumber partNumber, specialChars []specialChar) bool {
	return slices.ContainsFunc(specialChars, func(specialChar specialChar) bool {
		return contains(partNumber.topLeft, partNumber.bottomRight, specialChar.position)
	})
}

func contains(topLeft position, bottomRight position, pos position) bool {
	return pos.x >= topLeft.x && pos.x <= bottomRight.x && pos.y >= topLeft.y && pos.y <= bottomRight.y
}
