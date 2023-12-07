package day7_test

import (
	"aoc2023/day7"
	"reflect"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	type testArgs struct {
		cards   []string
		cardMap day7.CardMap
		result  string
	}
	tests := []testArgs{
		{
			cards:   []string{"3", "3", "2", "2", "2"},
			result:  day7.FullHouse,
			cardMap: day7.BuildCardMap([]string{"3", "3", "3", "2", "2"}),
		},
		{
			cards:   []string{"3", "3", "5", "6", "2"},
			result:  day7.Pair,
			cardMap: day7.BuildCardMap([]string{"3", "3", "5", "6", "2"}),
		},
		{
			cards:   []string{"3", "3", "3", "6", "2"},
			result:  day7.Three,
			cardMap: day7.BuildCardMap([]string{"3", "3", "3", "6", "2"}),
		},
		{
			cards:   []string{"3", "3", "3", "3", "2"},
			result:  day7.Four,
			cardMap: day7.BuildCardMap([]string{"3", "3", "3", "3", "2"}),
		},
		{
			cards:   []string{"3", "3", "3", "3", "3"},
			result:  day7.Five,
			cardMap: day7.BuildCardMap([]string{"3", "3", "3", "3", "3"}),
		},
		{
			cards:   []string{"2", "2", "3", "3", "4"},
			result:  day7.TwoPair,
			cardMap: day7.BuildCardMap([]string{"2", "2", "3", "3", "4"}),
		},
		{
			cards:   []string{"2", "5", "3", "6", "4"},
			result:  day7.High,
			cardMap: day7.BuildCardMap([]string{"2", "5", "3", "6", "4"}),
		},
	}

	for _, test := range tests {
		result := day7.CalculateScore(test.cardMap, test.cards)
		if result != test.result {
			t.Errorf("CalculateScore(%v, %v) = %v; wanted %v", test.cardMap, test.cards, result, test.result)
		}
	}
}

func TestCalculateScore2(t *testing.T) {
	type testArgs struct {
		cards   []string
		cardMap day7.CardMap
		result  string
	}
	tests := []testArgs{
		{
			cards:   []string{"3", "2", "T", "3", "K"},
			result:  day7.Pair,
			cardMap: day7.BuildCardMap2([]string{"3", "2", "T", "3", "K"}),
		},
		{
			cards:   []string{"T", "5", "5", "J", "5"},
			result:  day7.Four,
			cardMap: day7.BuildCardMap2([]string{"T", "5", "5", "J", "5"}),
		},
		{
			cards:   []string{"K", "K", "6", "7", "7"},
			result:  day7.TwoPair,
			cardMap: day7.BuildCardMap2([]string{"K", "K", "6", "7", "7"}),
		},
		{
			cards:   []string{"K", "T", "J", "J", "T"},
			result:  day7.Four,
			cardMap: day7.BuildCardMap2([]string{"K", "T", "J", "J", "T"}),
		},
		{
			cards:   []string{"Q", "Q", "Q", "J", "A"},
			result:  day7.Four,
			cardMap: day7.BuildCardMap2([]string{"Q", "Q", "Q", "J", "A"}),
		},
	}

	for _, test := range tests {
		result := day7.CalculateScore(test.cardMap, test.cards)
		if result != test.result {
			t.Errorf("CalculateScore(%v, %v) = %v; wanted %v", test.cardMap, test.cards, result, test.result)
		}
	}
}

// func TestSortHands(t *testing.T) {
// 	type testArgs struct {
// 		a      day7.Hand
// 		b      day7.Hand
// 		result bool
// 	}
// 	tests := []testArgs{
// 		{a: day7.BuildHand("QT539 20", 0), b: day7.BuildHand("K2351 20", 0), result: true},
// 		{a: day7.BuildHand("QQ539 20", 0), b: day7.BuildHand("KK351 20", 0), result: true},
// 		{a: day7.BuildHand("9Q539 20", 0), b: day7.BuildHand("9K351 20", 0), result: true},
// 		{a: day7.BuildHand("AQ539 20", 0), b: day7.BuildHand("9K351 20", 0), result: false},
// 		{a: day7.BuildHand("33355 20", 0), b: day7.BuildHand("2244K 20", 0), result: false},
// 	}

// 	for _, test := range tests {
// 		result := day7.SortByHighestCard(test.a, test.b)
// 		if result != test.result {
// 			t.Errorf("SortByHightestCard(%v, %v) = %v; wanted %v", test.a, test.b, result, test.result)
// 		}
// 	}
// }

func TestSumCardWinnings(t *testing.T) {
	result := day7.SumCardWinnings("test-input.txt", false)
	expected := 6440

	if result != expected {
		t.Errorf("SumCardWinnings() = %v; wanted %v", result, expected)
	}
}

func TestBuildCardMap2(t *testing.T) {
	type testArgs struct {
		cards  []string
		result day7.CardMap
	}
	tests := []testArgs{
		{
			cards:  []string{"3", "2", "T", "3", "K"},
			result: day7.CardMap{"3": 2, "2": 1, "T": 1, "K": 1},
		},
		{
			cards:  []string{"T", "5", "5", "J", "5"},
			result: day7.CardMap{"5": 4, "T": 1},
		},
		{
			cards:  []string{"K", "K", "6", "7", "7"},
			result: day7.CardMap{"K": 2, "7": 2, "6": 1},
		},
		{
			cards:  []string{"K", "T", "J", "J", "T"},
			result: day7.CardMap{"K": 1, "T": 4},
		},
		{
			cards:  []string{"Q", "Q", "Q", "J", "A"},
			result: day7.CardMap{"Q": 4, "A": 1},
		},
		{
			cards:  []string{"Q", "J", "3", "2", "5"},
			result: day7.CardMap{"Q": 2, "3": 1, "2": 1, "5": 1},
		},
		{
			cards:  []string{"Q", "Q", "3", "3", "J"},
			result: day7.CardMap{"Q": 3, "3": 2},
		},
		{
			cards:  []string{"Q", "J", "J", "J", "J"},
			result: day7.CardMap{"Q": 5},
		},
		{
			cards:  []string{"J", "J", "J", "J", "J"},
			result: day7.CardMap{"A": 5},
		},
		{
			cards:  []string{"3", "3", "J", "J", "J"},
			result: day7.CardMap{"3": 5},
		},
		{
			cards:  []string{"3", "3", "3", "J", "J"},
			result: day7.CardMap{"3": 5},
		},
	}

	for _, test := range tests {
		result := day7.BuildCardMap2(test.cards)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("BuildCardMap2(%v) = %v; wanted %v", test.cards, result, test.result)
		}
	}
}
