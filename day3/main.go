package main

import (
	"fmt"
	"util"
)

func main() {
	part1 := part1Parser{}
	util.ProcessFile("input.txt", &part1)
	fmt.Println("Day 3 part 1:", part1.calculateResult())

	part2 := part2Parser{}
	util.ProcessFile("input.txt", &part2)
	fmt.Println("Day 3 part 2:", part2.calculateResult())
}
