package main

import "testing"

func TestGetPartNumbersWithNoMatches(t *testing.T) {
	partNumbers := getPartNumbers("", 0)
	if len(partNumbers) > 0 {
		t.Errorf("getPartNumbers(\"\", 0) = %v, want empty", partNumbers)
	}

	partNumbers = getPartNumbers(".................", 0)
	if len(partNumbers) > 0 {
		t.Errorf("getPartNumbers(\"\", 0) = %v, want empty", partNumbers)
	}

	partNumbers = getPartNumbers("..#.@..$......&......", 0)
	if len(partNumbers) > 0 {
		t.Errorf("getPartNumbers(\"\", 0) = %v, want empty", partNumbers)
	}
}

func TestGetPartNumbersWithOneMatch(t *testing.T) {
	partNumbers := getPartNumbers("1..........", 0)
	if len(partNumbers) != 1 {
		t.Error("len(partNumbers) != 1")
	}

	partNumbers = getPartNumbers("..........1", 0)
	if len(partNumbers) != 1 {
		t.Error("len(partNumbers) != 1")
	}

	partNumbers = getPartNumbers(".....1.....", 0)
	if len(partNumbers) != 1 {
		t.Error("len(partNumbers) != 1")
	}
}

func TestGetPartNumbersWithMultipleMatches(t *testing.T) {
	partNumbers := getPartNumbers("1..........1", 0)
	if len(partNumbers) != 2 {
		t.Error("len(partNumbers) != 2")
	}

	partNumbers = getPartNumbers("1...1.....1", 0)
	if len(partNumbers) != 3 {
		t.Error("len(partNumbers) != 3")
	}

	partNumbers = getPartNumbers("1...1..1...1..1", 0)
	if len(partNumbers) != 5 {
		t.Error("len(partNumbers) != 5")
	}
}

func TestGetPartNumbersWithMultipleParts(t *testing.T) {
	partNumbers := getPartNumbers("1.....@.....2", 0)
	if len(partNumbers) != 2 {
		t.Error("len(partNumbers) != 2")
	}

	partNumbers = getPartNumbers("1...2.$....1", 0)
	if len(partNumbers) != 3 {
		t.Error("len(partNumbers) != 3")
	}

	partNumbers = getPartNumbers("1...2$..1...4..@3", 0)
	if len(partNumbers) != 5 {
		t.Error("len(partNumbers) != 5")
	}
}
