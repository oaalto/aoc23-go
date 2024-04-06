package util

import "bufio"

type NumberParser interface {
	Parse(string)
}

func Calculate(input string, numberParser NumberParser) {
	file := OpenFile(input)
	defer CloseFile(file)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		numberParser.Parse(fileScanner.Text())
	}
}
