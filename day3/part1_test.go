package main

import (
	"testing"
	"util"
)

func TestPart1TestInput(t *testing.T) {
	part1 := part1Parser{}
	util.ProcessFile("test_input.txt", &part1)

	if part1.calculateResult() != 4361 {
		t.Errorf("calculateResult() = %d, want %d", part1.calculateResult(), 4361)
	}
}

func TestPart1(t *testing.T) {
	part1 := part1Parser{}
	util.ProcessFile("input.txt", &part1)

	if part1.calculateResult() != 553825 {
		t.Errorf("calculateResult() = %d, want %d", part1.calculateResult(), 553825)
	}
}
