package main

import (
	"strings"
	"unicode"
)

type part1Parser int

func (p *part1Parser) Parse(line string) {
	firstIndex := strings.IndexFunc(line, unicode.IsDigit)
	if firstIndex == -1 {
		return
	}

	lastIndex := strings.LastIndexFunc(line, unicode.IsDigit)
	if lastIndex == -1 {
		return
	}

	*p += part1Parser(int(line[firstIndex]-'0')*10 + int(line[lastIndex]-'0'))
}
