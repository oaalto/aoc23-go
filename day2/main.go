package main

import (
	"fmt"
	"util"
)

func main() {
	var part1 part1Parser
	util.Calculate("input.txt", &part1)
	fmt.Println("Day 2 part 1:", part1)

	var part2 part2Parser
	util.Calculate("input.txt", &part2)
	fmt.Println("Day 2 part 2:", part2)
}

//Day 2 part 1: 1931
