package main

type part2Parser struct {
	curLine      int
	partNumbers  []partNumber
	specialChars []specialChar
}

func (parser *part2Parser) Parse(line string) {
	parser.partNumbers = append(parser.partNumbers, getPartNumbers(line, parser.curLine)...)
	parser.specialChars = append(parser.specialChars, getSpecialChars(line, parser.curLine)...)

	parser.curLine++
}

func (parser *part2Parser) calculateResult() int {
	result := 0
	for _, specialChar := range parser.specialChars {
		if gears := getGears(specialChar, parser.partNumbers); len(gears) == 2 {
			result += gears[0].number * gears[1].number
		}
	}
	return result
}

func getGears(specialChar specialChar, partNumbers []partNumber) []partNumber {
	gears := make([]partNumber, 0)
	for _, partNumber := range partNumbers {
		if specialChar.position.isInside(partNumber.topLeft, partNumber.bottomRight) {
			gears = append(gears, partNumber)
		}
	}
	return gears
}
