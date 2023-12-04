package day4

import (
	"aoc2023/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Id      int
	Winners []int
	Numbers []int
	Count   int
}

func parseCard(cardString string, _ int) Card {
	parts := strings.Split(cardString, ": ")
	idPart := parts[0]
	idString := strings.Replace(idPart, "Card ", "", 1)
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Printf("Failed to parse id %v\n", idString)
	}

	numbersParts := strings.Split(parts[1], " | ")
	winnerStrings := strings.Split(numbersParts[0], " ")
	winners := helpers.Map(winnerStrings, func(s string, _ int) int {
		if winner, err := strconv.Atoi(s); err == nil {
			return winner
		}
		fmt.Printf("failed to parse winning number %v\n", s)
		return 0
	})

	numberStrings := strings.Split(numbersParts[1], " ")
	numbers := helpers.Map(numberStrings, func(s string, _ int) int {
		if winner, err := strconv.Atoi(s); err == nil {
			return winner
		}
		fmt.Printf("failed to parse number %v\n", s)
		return 0
	})
	return Card{Id: id, Winners: winners, Numbers: numbers, Count: 1}
}

func parseInput(input string) []Card {
	re := regexp.MustCompile(` +`)
	cleanedInput := re.ReplaceAllString(input, " ")
	cardStrings := strings.Split(cleanedInput, "\n")

	return helpers.Map(cardStrings, parseCard)
}

func SumWinningCardPoints(input string) int {
	text := helpers.ReadInput(input)
	cards := parseInput(text)

	winnerScores := helpers.Map(cards, func(card Card, _ int) int {
		return helpers.Reduce(card.Numbers, func(result int, number int, _ int) int {
			_, err := helpers.FindIndex(card.Winners, func(winner int) bool {
				return winner == number
			})
			if err != nil {
				return result
			}
			if result == 0 {
				return 1
			}
			return result * 2
		}, 0)
	})
	return helpers.Sum(winnerScores)
}

func SumTotalCards(input string) int {
	text := helpers.ReadInput(input)

	cards := parseInput(text)
	for i, card := range cards {
		winningNums := []int{}
		for _, number := range card.Numbers {
			_, err := helpers.FindIndex(card.Winners, func(winner int) bool {
				return winner == number
			})
			if err != nil {
				continue
			}
			winningNums = append(winningNums, number)
		}

		for j := i + 1; j-i <= len(winningNums); j = j + 1 {
			cards[j].Count = cards[j].Count + card.Count
		}
	}

	return helpers.Reduce(cards, func(result int, card Card, _ int) int {
		return result + card.Count
	}, 0)
}

func Part1() string {
	return fmt.Sprintf("%v", SumWinningCardPoints("day4/input.txt"))
}

func Part2() string {
	return fmt.Sprintf("%v", SumTotalCards("day4/input.txt"))
}
