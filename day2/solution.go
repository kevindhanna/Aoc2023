package day2

import (
	"aoc2023/helpers"
	"errors"
	"fmt"
)

func find(set Round, wanted Cube) (Cube, error) {
	for _, cube := range set {
		if cube.Colour == wanted.Colour {
			return cube, nil
		}
	}
	return Cube{}, errors.New(fmt.Sprintf("No match in set %v for wanted %v", set, wanted))
}

func IsPossible(game Game, set Round) bool {
	for _, round := range game.Rounds {
		for _, cube := range round {
			setEquivalent, err := find(set, cube)
			if err != nil || cube.Count > setEquivalent.Count {
				return false
			}
		}
	}

	return true
}

var part1Set = []Cube{
	{12, "red"},
	{13, "green"},
	{14, "blue"},
}

func SumPossibleGameIds(input string) int {
	gameStrings := helpers.ReadInputToLines(input)
	games := parseGames(gameStrings)

	possibleGames := helpers.Filter(games, func(game Game) bool {
		return IsPossible(game, part1Set)
	})
	return helpers.Reduce(possibleGames, func(total int, game Game) int {
		return total + game.Id
	}, 0)
}

func Part1() string {
	result := SumPossibleGameIds("day2/input.txt")
	return fmt.Sprint(result)
}

func CalculatePower(game Game) int {
	maxCubes := helpers.Reduce(game.Rounds, func(result []Cube, round Round) []Cube {
		for _, cube := range round {
			i, err := helpers.FindIndex(result, func(max Cube) bool {
				return max.Colour == cube.Colour
			})
			if err != nil {
				result = append(result, cube)
			} else {
				max := result[i]
				if max.Count < cube.Count {
					result[i] = cube
				}
			}
		}
		return result
	}, []Cube{})

	counts := helpers.Map(maxCubes, func(cube Cube) int {
		return cube.Count
	})
	return helpers.Reduce(counts, func(result int, count int) int {
		return result * count
	}, 1)
}

func SumGamePowers(input string) int {
	gameStrings := helpers.ReadInputToLines(input)
	games := parseGames(gameStrings)

	return helpers.Sum(helpers.Map(games, func(game Game) int {
		return CalculatePower(game)
	}))
}

func Part2() string {
	result := SumGamePowers("day2/input.txt")
	return fmt.Sprint(result)
}
