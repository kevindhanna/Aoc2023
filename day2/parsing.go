package day2

import (
	"aoc2023/helpers"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Cube struct {
	Count  int
	Colour string
}

type Round []Cube

type Game struct {
	Id     int
	Rounds []Round
}

func ParseCube(resultString string, _ int) (Cube, error) {
	resultTuple := strings.Split(resultString, " ")
	if len(resultTuple) != 2 {
		return Cube{}, errors.New(fmt.Sprintf("Failed to parse game result: %v", resultString))
	}

	count, err := strconv.Atoi(resultTuple[0])
	if err != nil {
		return Cube{}, errors.New(fmt.Sprintf("Failed to parse game result count: %v", resultTuple[1]))
	}

	return Cube{Colour: resultTuple[1], Count: count}, nil
}

func ParseRound(roundPart string, _ int) (Round, error) {
	return helpers.SplitMap(roundPart, ", ", ParseCube)
}

func ParseRounds(roundsPart string) ([]Round, error) {
	return helpers.SplitMap(roundsPart, "; ", ParseRound)
}

func parseId(idPart string) (int, error) {
	idString := strings.Replace(idPart, "Game ", "", 1)
	return strconv.Atoi(idString)
}

func buildGame(gameString string) (Game, error) {
	gameSlice := strings.Split(gameString, ": ")
	idPart, resultsPart := gameSlice[0], gameSlice[1]

	if len(gameSlice) > 2 {
		return Game{}, errors.New(fmt.Sprintf("Failed to parse game: %v", gameString))
	}

	id, err := parseId(idPart)
	if err != nil {
		return Game{}, err
	}

	results, err := ParseRounds(resultsPart)

	return Game{Id: id, Rounds: results}, nil
}

func parseGames(gameStrings []string) []Game {
	games := []Game{}
	for _, gameString := range gameStrings {
		g, err := buildGame(gameString)
		if err != nil {
			fmt.Printf("Invalid game string: %v", gameString)
		}
		games = append(games, g)
	}
	return games
}
