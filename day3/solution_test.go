package day3_test

import (
	"aoc2023/day3"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGetSurroundingCoords(t *testing.T) {
	type testArgs struct {
		coord  day3.Coord
		maxX   int
		maxY   int
		result []day3.Coord
	}
	tests := []testArgs{
		{
			coord: day3.Coord{5, 5},
			maxX:  10,
			maxY:  10,
			result: []day3.Coord{
				{4, 4},
				{5, 4},
				{6, 4},
				{4, 5},
				{6, 5},
				{4, 6},
				{5, 6},
				{6, 6},
			},
		},
		{
			coord: day3.Coord{0, 0},
			maxX:  10,
			maxY:  10,
			result: []day3.Coord{
				{1, 0},
				{0, 1},
				{1, 1},
			},
		},
		{
			coord: day3.Coord{10, 10},
			maxX:  10,
			maxY:  10,
			result: []day3.Coord{
				{9, 9},
				{10, 9},
				{9, 10},
			},
		},
	}
	for _, test := range tests {
		result := day3.GetSurroundingCoords(test.coord, test.maxX, test.maxY)
		if !reflect.DeepEqual(result, test.result) {
			t.Error(fmt.Sprintf("GetSurroundingCoords(%v, %v, %v) \n= %v; \n! %v\n", test.coord, test.maxX, test.maxY, result, test.result))
		}
	}
}

func TestGetRestOfNumber(t *testing.T) {
	type testArgs struct {
		start      int
		line       string
		numString  string
		usedPoints []int
	}

	tests := []testArgs{
		{start: 5, line: "....1234....", numString: "1234", usedPoints: []int{4, 5, 6, 7}},
		{start: 3, line: "1234....", numString: "1234", usedPoints: []int{0, 1, 2, 3}},
		{start: 5, line: "....1234", numString: "1234", usedPoints: []int{4, 5, 6, 7}},
	}

	for _, test := range tests {
		lineSlice := strings.Split(test.line, "")
		numString, usedPoints := day3.GetRestOfNumber(test.start, lineSlice)
		if numString != test.numString {
			t.Error(fmt.Printf("GetRestOfNumber(%v, %v) = %v; wanted %v\n", test.start, test.line, numString, test.numString))
		}

		sort.Slice(usedPoints, func(i, j int) bool {
			return usedPoints[i] < usedPoints[j]
		})
		if !reflect.DeepEqual(usedPoints, test.usedPoints) {
			t.Error(fmt.Printf("GetRestOfNumber(%v, %v) = %v; wanted %v\n", test.start, test.line, usedPoints, test.usedPoints))
		}

	}
}

func TestSumPartNumbers(t *testing.T) {
	result := day3.SumPartNumbers("test-input.txt")
	if result != 4361 {
		t.Error(fmt.Printf("SumPartNumbers() = %v; wanted %v", result, 4361))
	}
}

func TestSumGearRatios(t *testing.T) {
	result := day3.SumGearRatios("test-input.txt")
	if result != 467835 {
		t.Error(fmt.Printf("SumGearRatios() = %v; wanted %v", result, 467835))
	}
}
