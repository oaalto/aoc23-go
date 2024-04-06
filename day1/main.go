package main

import (
	"advent-of-code-go/util"
	"bufio"
	"fmt"
)

type numberParser interface {
	parse(string)
}

func main() {
	var part1 part1Parser = 0
	calculate("./day1/input.txt", &part1)
	fmt.Println("Day 1 part 1: ", part1)

	var part2 part2Parser = 0
	calculate("./day1/input.txt", &part2)
	fmt.Println("Day 1 part 2: ", part2)
}

func calculate(input string, numberParser numberParser) {
	file := util.OpenFile(input)
	defer util.CloseFile(file)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		numberParser.parse(fileScanner.Text())
	}
}
