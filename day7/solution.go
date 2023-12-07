package day7

import (
	"aoc2023/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	High      = "H"
	Pair      = "P"
	TwoPair   = "2P"
	Three     = "T"
	FullHouse = "FH"
	Four      = "4"
	Five      = "5"
)

var cardWeights1 = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var cardWeights2 = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
var handWeights = []string{High, Pair, TwoPair, Three, FullHouse, Four, Five}

type CardMap map[string]int
type Hand struct {
	Cards     []string
	CardMap   CardMap
	ScoreName string
	ScoreVal  int
	Bid       int
}
type ByScore []Hand
type ByScore2 []Hand

func SortByHighestCard(a, b Hand) bool {
	for i, aCard := range a.Cards {
		bCard := b.Cards[i]
		if aCard == bCard {
			continue
		}
		aWeight, err := helpers.FindIndexComparable(cardWeights1, aCard)
		if err != nil {
			panic(fmt.Sprintf("Cound't find weight index for card %v... which is weird", aCard))
		}
		bWeight, err := helpers.FindIndexComparable(cardWeights1, bCard)
		if err != nil {
			panic(fmt.Sprintf("Cound't find weight index for card %v... which is weird", bCard))
		}
		return aWeight < bWeight
	}
	return false
}
func SortByHighestCard2(a, b Hand) bool {
	for i, aCard := range a.Cards {
		bCard := b.Cards[i]
		if aCard == bCard {
			continue
		}
		aWeight, err := helpers.FindIndexComparable(cardWeights2, aCard)
		if err != nil {
			panic(fmt.Sprintf("Cound't find weight index for card %v... which is weird", aCard))
		}
		bWeight, err := helpers.FindIndexComparable(cardWeights2, bCard)
		if err != nil {
			panic(fmt.Sprintf("Cound't find weight index for card %v... which is weird", bCard))
		}
		return aWeight < bWeight
	}
	return false
}

func (hands ByScore) Less(i, j int) bool {
	result := hands[i].ScoreVal < hands[j].ScoreVal
	if hands[i].ScoreName == hands[j].ScoreName {
		result = SortByHighestCard(hands[i], hands[j])
	}

	return result
}

func (hands ByScore2) Less(i, j int) bool {
	result := hands[i].ScoreVal < hands[j].ScoreVal
	if hands[i].ScoreName == hands[j].ScoreName {
		result = SortByHighestCard2(hands[i], hands[j])
	}

	return result
}
func (hands ByScore) Len() int {
	return len(hands)
}
func (hands ByScore) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}
func (hands ByScore2) Len() int {
	return len(hands)
}
func (hands ByScore2) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func CalculateScore(cardMap CardMap, cards []string) string {
	keys := helpers.MapKeys(cardMap)

	scores := []string{}
	for _, k := range keys {
		count := cardMap[k]
		switch count {
		case 1:
			continue
		case 2:
			scores = append(scores, Pair)
		case 3:
			scores = append(scores, Three)
		case 4:
			scores = append(scores, Four)
		case 5:
			scores = append(scores, Five)
		}
	}

	if len(scores) == 2 {
		if scores[0] == Pair && scores[1] == Pair {
			return TwoPair
		}
		return FullHouse
	} else if len(scores) == 0 {
		return High
	}
	return scores[0]
}

func BuildCardMap(cards []string) CardMap {
	return helpers.Reduce(cards, func(cardMap CardMap, card string, _ int) CardMap {
		cardMap[card] = cardMap[card] + 1
		return cardMap
	}, CardMap{})
}

func buildHandFactory(mapBuilder func(cards []string) CardMap) func(handString string, _ int) Hand {
	return func(handString string, _ int) Hand {
		parts := strings.Split(handString, " ")
		cards := strings.Split(parts[0], "")
		cardMap := mapBuilder(cards)
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to parse bid %v, error: %v", parts[1], err)
		}
		scoreName := CalculateScore(cardMap, cards)
		scoreVal, err := helpers.FindIndexComparable(handWeights, scoreName)
		if err != nil {
			panic(fmt.Sprintf("Cound't find card hand weight index for score %v... which is weird", scoreName))
		}

		return Hand{
			Cards:     cards,
			CardMap:   cardMap,
			ScoreName: scoreName,
			ScoreVal:  scoreVal,
			Bid:       bid,
		}
	}
}

func SumCardWinnings(input string, part2 bool) int {
	var handBuilder func(handString string, _ int) Hand
	if part2 {
		handBuilder = buildHandFactory(BuildCardMap2)
	} else {
		handBuilder = buildHandFactory(BuildCardMap)
	}
	handStrings := helpers.ReadInputToLines(input)
	hands := helpers.Map(handStrings, handBuilder)
	if part2 {
		sort.Sort(ByScore2(hands))
	} else {
		sort.Sort(ByScore(hands))
	}

	return helpers.Reduce(hands, func(result int, hand Hand, i int) int {
		return result + (hand.Bid * (i + 1))
	}, 0)
}

func Part1() string {
	return fmt.Sprint(SumCardWinnings("day7/input.txt", false))
}

func BuildCardMap2(cards []string) CardMap {
	m := helpers.Reduce(cards, func(cardMap CardMap, card string, _ int) CardMap {
		cardMap[card] = cardMap[card] + 1
		return cardMap
	}, CardMap{})
	keys := helpers.MapKeys(m)
	sort.Slice(keys, func(i, j int) bool {
		iWeight, err := helpers.FindIndexComparable(cardWeights2, keys[i])
		if err != nil {
			panic(fmt.Sprintf("Cound't find weight index for card %v... which is weird", keys[i]))
		}
		jWeight, err := helpers.FindIndexComparable(cardWeights2, keys[j])
		if err != nil {
			panic(fmt.Sprintf("Cound't find weight index for card %v... which is weird", keys[j]))
		}
		return iWeight > jWeight
	})
	Js := m["J"]
	delete(m, "J")

	if Js == 5 {
		m["A"] = 5
	}
	if Js > 0 {
		score := CalculateScore(m, cards)
		switch score {
		case High:
			m[keys[0]] = m[keys[0]] + Js
		case Pair:
			fallthrough
		case TwoPair:
			for _, key := range keys {
				if m[key] == 2 {
					m[key] = m[key] + Js
					break
				}
			}
		case Three:
			for _, key := range keys {
				if m[key] == 3 {
					m[key] = m[key] + Js
					break
				}
			}
		case Four:
			for _, key := range keys {
				if m[key] == 4 {
					m[key] = m[key] + Js
					break
				}
			}
		}
	}
	return m
}

func Part2() string {
	return fmt.Sprint(SumCardWinnings("day7/input.txt", true))
}
