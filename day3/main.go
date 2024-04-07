package main

import (
	"fmt"
	"util"
)

func main() {
	part1 := part1Parser{}
	util.ProcessFile("input.txt", &part1)
	fmt.Println("Day 3 part 1:", part1.calculateResult()) // 553825
}
