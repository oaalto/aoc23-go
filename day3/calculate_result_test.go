package main

import (
	"strings"
	"testing"
)

var input = `1..@..2.
.$..3..@
...4....
......10
.4....&.`

func TestCalculateResult(t *testing.T) {
	part1Parser := part1Parser{}

	for _, line := range strings.Split(input, "\n") {
		part1Parser.Parse(line)
	}

	result := part1Parser.calculateResult()
	if result != 16 {
		t.Error("Expected 16 results, got ", result)
	}
}
