package main

import (
	"regexp"
	"util"
)

var numberRegexp = regexp.MustCompile("[0-9]+")

type partNumber struct {
	number      int
	topLeft     position
	bottomRight position
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
