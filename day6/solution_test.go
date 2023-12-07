package day6_test

import (
	"aoc2023/day6"
	"testing"
)

func TestGetMinHoldTime(t *testing.T) {
	type testArgs struct {
		time           int
		distanceRecord int
		result         int
	}
	tests := []testArgs{
		{time: 7, distanceRecord: 9, result: 2},
		{time: 15, distanceRecord: 40, result: 4},
		{time: 30, distanceRecord: 200, result: 11},
	}

	for _, test := range tests {
		result := day6.GetMinHoldTime(test.time, test.distanceRecord)
		if result != test.result {
			t.Errorf("GetMinHoldTime(%v, %v) = %v; wanted %v", test.time, test.distanceRecord, result, test.result)
		}
	}
}

func TestGetMaxHoldTime(t *testing.T) {
	type testArgs struct {
		time           int
		distanceRecord int
		min            int
		result         int
	}
	tests := []testArgs{
		{time: 7, distanceRecord: 9, min: 2, result: 5},
		{time: 15, distanceRecord: 40, min: 4, result: 11},
		{time: 30, distanceRecord: 200, min: 11, result: 19},
	}

	for _, test := range tests {
		result := day6.GetMaxHoldTime(test.time, test.distanceRecord, test.min)
		if result != test.result {
			t.Errorf("GetMaxHoldTime(%v, %v) = %v; wanted %v", test.time, test.distanceRecord, result, test.result)
		}
	}
}

func TestGetWinStrategyCount(t *testing.T) {
	type testArgs struct {
		race   day6.Race
		result int
	}
	tests := []testArgs{
		{race: day6.Race{Time: 7, DistanceRecord: 9}, result: 4},
		{race: day6.Race{Time: 15, DistanceRecord: 40}, result: 8},
		{race: day6.Race{Time: 30, DistanceRecord: 200}, result: 9},
	}

	for _, test := range tests {
		result := day6.GetWinStrategyCount(test.race)
		if result != test.result {
			t.Errorf("GetWinStrategyCount(%v) = %v; wanted %v", test.race, result, test.result)
		}
	}
}

func TestGetWinStrategyProduct(t *testing.T) {
	result := day6.GetWinStrategyProduct("test-input.txt")
	expected := 288
	if result != expected {
		t.Errorf("GetWinStrategyProduct() = %v; wanted %v", result, expected)
	}
}

func TestGetWinStrategyCountOfCombinedRace(t *testing.T) {
	result := day6.GetWinStrategyCountOfCombinedRace("test-input.txt")
	expected := 71503
	if result != expected {
		t.Errorf("GetWinStrategyCountOfCombinedRaces() = %v; wanted %v", result, expected)
	}
}
