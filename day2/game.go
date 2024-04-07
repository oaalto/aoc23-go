package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"util"
)

type game struct {
	red   int
	green int
	blue  int
}

func parseGame(line string) game {
	split := strings.Split(line, ",")
	game := game{}

	for _, gamePart := range split {
		gamePart := strings.TrimSpace(gamePart)
		parts := strings.Split(gamePart, " ")

		switch parts[1] {
		case "red":
			game.red = util.MustParseInt(parts[0])
		case "green":
			game.green = util.MustParseInt(parts[0])
		case "blue":
			game.blue = util.MustParseInt(parts[0])
		}
	}

	return game
}

func parseGames(line string) []game {
	split := strings.Split(line, ";")
	games := make([]game, 0)
	for _, game := range split {
		games = append(games, parseGame(game))
	}

	return games
}

var gameIdRegexp = regexp.MustCompile(`^Game ([0-9]+)`)

func parseGameId(line string) int {
	matches := gameIdRegexp.FindStringSubmatch(line)
	if len(matches) != 2 {
		log.Fatal("Invalid game id line:", line)
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal("Invalid game id line:", line)
	}

	return id
}
