package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type numberIndex struct {
	index  int
	number string
}

type part2Parser int

func (p *part2Parser) parse(line string) {
	numberIndices := make([]numberIndex, 0)

	if index := strings.IndexFunc(line, unicode.IsDigit); index > -1 {
		numberIndices = append(numberIndices, numberIndex{index: index, number: string(line[index])})
	}

	if index := strings.LastIndexFunc(line, unicode.IsDigit); index > -1 {
		numberIndices = append(numberIndices, numberIndex{index: index, number: string(line[index])})
	}

	wordIndices := findWordIndices(line)

	numberIndices = append(numberIndices, wordIndices...)

	slices.SortFunc(numberIndices, func(a numberIndex, b numberIndex) int {
		return a.index - b.index
	})

	number, err := strconv.Atoi(numberIndices[0].number + numberIndices[len(numberIndices)-1].number)
	if err != nil {
		log.Fatal(numberIndices[0], "not a number")
	}

	*p += part2Parser(number)
}

func findWordIndices(line string) []numberIndex {
	numberIndices := make([]numberIndex, 0)
	for word, number := range numberMap {
		if index := strings.Index(line, word); index > -1 {
			numberIndices = append(numberIndices, numberIndex{index: index, number: number})
		}

		if index := strings.LastIndex(line, word); index > -1 {
			numberIndices = append(numberIndices, numberIndex{index: index, number: number})
		}
	}

	return numberIndices
}
