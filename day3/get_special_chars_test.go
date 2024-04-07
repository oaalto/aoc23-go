package main

import "testing"

func TestGetSpecialCharsWithNoMatches(t *testing.T) {
	specialChars := getSpecialChars("", 0)
	if len(specialChars) > 0 {
		t.Error("Expected no special chars, got ", specialChars)
	}

	specialChars = getSpecialChars(".................", 0)
	if len(specialChars) > 0 {
		t.Error("Expected no special chars, got ", specialChars)
	}

	specialChars = getSpecialChars("2..3..45..657....4534..44..2...", 0)
	if len(specialChars) > 0 {
		t.Error("Expected no special chars, got ", specialChars)
	}

	specialChars = getSpecialChars("123456789", 0)
	if len(specialChars) > 0 {
		t.Error("Expected no special chars, got ", specialChars)
	}
}

func TestGetSpecialCharsWithOneMatch(t *testing.T) {
	specialChars := getSpecialChars(".43...@....", 0)
	if len(specialChars) != 1 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 6, t)

	specialChars = getSpecialChars("@...43.", 0)
	if len(specialChars) != 1 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)

	specialChars = getSpecialChars("....23.@", 0)
	if len(specialChars) != 1 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 7, t)
}

func TestGetSpecialCharsWithMultipleMatches(t *testing.T) {
	specialChars := getSpecialChars("@....@...4.", 0)
	if len(specialChars) != 2 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)
	assertSpecialChar(specialChars[1], 5, t)

	specialChars = getSpecialChars("@..3..@", 0)
	if len(specialChars) != 2 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)
	assertSpecialChar(specialChars[1], 6, t)

	specialChars = getSpecialChars("@.3..@..@", 0)
	if len(specialChars) != 3 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)
	assertSpecialChar(specialChars[1], 5, t)
	assertSpecialChar(specialChars[2], 8, t)
}

func TestGetSpecialCharsWithMultipleChars(t *testing.T) {
	specialChars := getSpecialChars("@....$...34.", 0)
	if len(specialChars) != 2 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)
	assertSpecialChar(specialChars[1], 5, t)

	specialChars = getSpecialChars("@...$565", 0)
	if len(specialChars) != 2 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)
	assertSpecialChar(specialChars[1], 4, t)

	specialChars = getSpecialChars("@..324234.$..@434", 0)
	if len(specialChars) != 3 {
		t.Error("Expected no special chars, got ", specialChars)
	}
	assertSpecialChar(specialChars[0], 0, t)
	assertSpecialChar(specialChars[1], 10, t)
	assertSpecialChar(specialChars[2], 13, t)
}

func assertSpecialChar(specialChar specialChar, posX int, t *testing.T) {
	if specialChar.position.x != posX {
		t.Error("not matching positionX, expected", posX, "was", specialChar.position.x)
	}
}
