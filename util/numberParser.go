package util

import (
	"bufio"
	"log"
	"strconv"
)

type NumberParser interface {
	Parse(string)
}

func ProcessFile(input string, numberParser NumberParser) {
	file := OpenFile(input)
	defer CloseFile(file)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		numberParser.Parse(fileScanner.Text())
	}
}

func MustParseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
