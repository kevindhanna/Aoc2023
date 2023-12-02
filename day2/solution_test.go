package day2_test

import (
	"aoc2023/day2"
	"fmt"
	"reflect"
	"testing"
)

func TestParseCube(t *testing.T) {
	type testArgs struct {
		input  string
		result day2.Cube
	}
	tests := []testArgs{
		{input: "2 red", result: day2.Cube{2, "red"}},
	}

	for _, test := range tests {
		result, err := day2.ParseCube(test.input)
		if err != nil {
			t.Error(fmt.Sprintf("parseCube(%v) = %v; wanted %v", test.input, err, test.result))
		}
		if result != test.result {
			t.Error(fmt.Sprintf("parseCube(%v) = %v; wanted %v", test.input, result, test.result))
		}
	}
}

func TestParseRound(t *testing.T) {
	type testArgs struct {
		input  string
		result day2.Round
	}
	tests := []testArgs{
		{
			input: "1 red, 2 green, 6 blue",
			result: day2.Round{
				{1, "red"},
				{2, "green"},
				{6, "blue"},
			},
		},
	}

	for _, test := range tests {
		result, err := day2.ParseRound(test.input)
		if err != nil {
			t.Error(fmt.Sprintf("parseRound(%v) = %v; wanted %v", test.input, err, test.result))
		}
		if !reflect.DeepEqual(result, test.result) {
			t.Error(fmt.Sprintf("parseRound(%v) = %v; wanted %v", test.input, result, test.result))
		}
	}
}

func TestParseRounds(t *testing.T) {
	type testArgs struct {
		input  string
		result []day2.Round
	}
	tests := []testArgs{
		{
			input: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			result: []day2.Round{
				{
					{3, "blue"},
					{4, "red"},
				},
				{
					{1, "red"},
					{2, "green"},
					{6, "blue"},
				},
				{
					{2, "green"},
				},
			},
		},
	}

	for _, test := range tests {
		result, err := day2.ParseRounds(test.input)
		if err != nil {
			t.Error(fmt.Sprintf("parseRound(%v) = %v; wanted %v", test.input, err, test.result))
		}
		if !reflect.DeepEqual(result, test.result) {
			t.Error(fmt.Sprintf("parseRound(%v) = %v; got %v", test.input, result, test.result))
		}
	}
}

func TestIsPossible(t *testing.T) {
	type testArgs struct {
		set    day2.Round
		game   day2.Game
		result bool
	}

	tests := []testArgs{
		{
			set: day2.Round{
				{6, "blue"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{6, "blue"}},
				},
			},
			result: true,
		},
		{
			set: day2.Round{
				{6, "blue"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{7, "blue"}},
				},
			},
			result: false,
		},
		{
			set: day2.Round{
				{6, "blue"},
				{5, "red"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{6, "blue"}},
					{{5, "red"}},
				},
			},
			result: true,
		},
		{
			set: day2.Round{
				{6, "blue"},
				{5, "red"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{6, "blue"}},
					{{7, "red"}},
				},
			},
			result: false,
		},
		{
			set: day2.Round{
				{6, "blue"},
				{5, "red"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{6, "blue"}, {4, "red"}},
					{{2, "blue"}, {2, "red"}},
				},
			},
			result: true,
		},
		{
			set: day2.Round{
				{6, "blue"},
				{5, "red"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{6, "blue"}, {4, "red"}},
					{{2, "blue"}, {7, "red"}},
				},
			},
			result: false,
		},
		{
			set: day2.Round{
				{6, "blue"},
				{5, "red"},
			},
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{6, "blue"}, {4, "red"}},
					{{2, "blue"}, {2, "red"}},
					{{1, "green"}},
				},
			},
			result: false,
		},
	}

	for _, test := range tests {
		result := day2.IsPossible(test.game, test.set)
		if result != test.result {
			t.Error(fmt.Sprintf("parseCube(%v, %v) = %v; wanted %v", test.set, test.game, result, test.result))
		}
	}
}

func TestSumPossibleGameIds(t *testing.T) {
	input := "test-input.txt"
	result := day2.SumPossibleGameIds(input)
	if result != 8 {
		t.Error(fmt.Sprintf("Part1(%v) = %v; wanted %v", input, result, 8))
	}
}

func TestCalculatePower(t *testing.T) {
	type testArgs struct {
		game   day2.Game
		result int
	}

	tests := []testArgs{
		{
			result: 48,
			game: day2.Game{
				Id: 1,
				Rounds: []day2.Round{
					{{3, "blue"}, {4, "red"}},
					{{1, "red"}, {2, "green"}, {6, "blue"}},
					{{2, "green"}},
				},
			},
		},
	}
	for _, test := range tests {
		result := day2.CalculatePower(test.game)
		if result != test.result {
			t.Error(fmt.Sprintf("CalculatePower(%v) = %v; wanted %v", test.game, result, test.result))
		}
	}
}

func TestSumGamePowers(t *testing.T) {
	input := "test-input.txt"
	result := day2.SumGamePowers(input)

	if result != 2286 {
		t.Error(fmt.Sprintf("Part2(%v) = %v; wanted %v", input, result, 2286))
	}
}
