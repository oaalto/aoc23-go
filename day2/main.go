package main

import (
	"fmt"
	"util"
)

func main() {
	var part1 part1Parser
	util.ProcessFile("input.txt", &part1)
	fmt.Println("Day 2 part 1:", part1)

	var part2 part2Parser
	util.ProcessFile("input.txt", &part2)
	fmt.Println("Day 2 part 2:", part2)
}
