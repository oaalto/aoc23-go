package main

import (
	"testing"
)

var inputLine1 = "1..@..2."
var inputLine2 = ".$..3..@"
var inputLine3 = "...4...."
var inputLine4 = "......10"
var inputLine5 = ".4....&."

func TestParse(t *testing.T) {
	part1Parser := part1Parser{}

	part1Parser.Parse(inputLine1)

	if part1Parser.curLine != 1 {
		t.Errorf("curLine should be 1, was %d", part1Parser.curLine)
	}

	if len(part1Parser.partNumbers) != 2 {
		t.Errorf("partNumber count should be 2, was %d", len(part1Parser.partNumbers))
	}

	if len(part1Parser.specialChars) != 1 {
		t.Errorf("specialCharCount should be 1, was %d", len(part1Parser.specialChars))
	}

	part1Parser.Parse(inputLine2)

	if part1Parser.curLine != 2 {
		t.Errorf("curLine should be 2, was %d", part1Parser.curLine)
	}

	if len(part1Parser.partNumbers) != 3 {
		t.Errorf("partNumber count should be 3, was %d", len(part1Parser.partNumbers))
	}

	if len(part1Parser.specialChars) != 3 {
		t.Errorf("specialCharCount should be 3, was %d", len(part1Parser.specialChars))
	}

	part1Parser.Parse(inputLine3)

	if part1Parser.curLine != 3 {
		t.Errorf("curLine should be 3, was %d", part1Parser.curLine)
	}

	if len(part1Parser.partNumbers) != 4 {
		t.Errorf("partNumber count should be 4, was %d", len(part1Parser.partNumbers))
	}

	if len(part1Parser.specialChars) != 3 {
		t.Errorf("specialCharCount should be 3, was %d", len(part1Parser.specialChars))
	}

	part1Parser.Parse(inputLine4)

	if part1Parser.curLine != 4 {
		t.Errorf("curLine should be 4, was %d", part1Parser.curLine)
	}

	if len(part1Parser.partNumbers) != 5 {
		t.Errorf("partNumber count should be 5, was %d", len(part1Parser.partNumbers))
	}

	if len(part1Parser.specialChars) != 3 {
		t.Errorf("specialCharCount should be 3, was %d", len(part1Parser.specialChars))
	}

	part1Parser.Parse(inputLine5)

	if part1Parser.curLine != 5 {
		t.Errorf("curLine should be 5, was %d", part1Parser.curLine)
	}

	if len(part1Parser.partNumbers) != 6 {
		t.Errorf("partNumber count should be 6, was %d", len(part1Parser.partNumbers))
	}

	if len(part1Parser.specialChars) != 4 {
		t.Errorf("specialCharCount should be 4, was %d", len(part1Parser.specialChars))
	}
}
