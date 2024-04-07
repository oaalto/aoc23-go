package main

import "strings"

type part2Parser int

func (p *part2Parser) Parse(line string) {
	split := strings.Split(line, ":")
	games := parseGames(split[1])

	maxPossibleGame := getMaxPossibleGame(games)
	*p += part2Parser(maxPossibleGame.red * maxPossibleGame.green * maxPossibleGame.blue)
}

func getMaxPossibleGame(games []game) game {
	game := game{}
	for _, g := range games {
		game.red = max(game.red, g.red)
		game.green = max(game.green, g.green)
		game.blue = max(game.blue, g.blue)
	}

	return game
}
