package main

import (
	"aoc23-go/util"
	"fmt"
)

func main() {
	var part1 part1Parser
	util.Calculate("./day1/input.txt", &part1)
	fmt.Println("Day 1 part 1: ", part1)

	var part2 part2Parser
	util.Calculate("./day1/input.txt", &part2)
	fmt.Println("Day 1 part 2: ", part2)
}
