package main

import (
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type part1Parser int

func (p *part1Parser) Parse(line string) {
	split := strings.Split(line, ":")
	games := parseGames(split[1])

	if isPossible(games) {
		gameId := parseGameId(line)
		*p += part1Parser(gameId)
	}
}

func isPossible(games []game) bool {
	for _, g := range games {
		if g.red > maxRed || g.green > maxGreen || g.blue > maxBlue {
			return false
		}
	}

	return true
}
