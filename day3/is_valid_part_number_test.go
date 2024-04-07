package main

import "testing"

var notMatchingPartNumber = partNumber{number: 1, topLeft: position{x: -1, y: -1}, bottomRight: position{x: 1, y: 1}}
var matchingPartNumber = partNumber{number: 2, topLeft: position{x: 9, y: -1}, bottomRight: position{x: 11, y: 1}}

var specialCharDollar = specialChar{position: position{x: 10, y: 0}}
var specialCharEqual = specialChar{position: position{x: 0, y: 2}}

func TestIsValidPartNumberWithNoMatches(t *testing.T) {
	specialChars := []specialChar{specialCharDollar, specialCharEqual}
	if isValidPartNumber(notMatchingPartNumber, specialChars) {
		t.Error("Expected to not match anything")
	}
}

func TestIsValidPartNumberWithMatch(t *testing.T) {
	specialChars := []specialChar{specialCharDollar, specialCharEqual}
	if !isValidPartNumber(matchingPartNumber, specialChars) {
		t.Error("Expected to match")
	}
}
