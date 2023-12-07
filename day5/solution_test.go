package day5_test

import (
	"aoc2023/day5"
	"reflect"
	"testing"
)

func TestParseMapLine(t *testing.T) {
	result := day5.ParseMapLine("10 10 5", false)
	expected := day5.Map{Destination: 10, Source: 10, Length: 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseMapLine() = %v; wanted %v", result, expected)
	}
}

func TestGetRelationForMap(t *testing.T) {
	type testArgs struct {
		i      int
		m      day5.Map
		result int
		err    bool
	}
	tests := []testArgs{
		{i: 5, m: day5.Map{Source: 5, Destination: 10, Length: 5}, result: 10},
		{i: 10, m: day5.Map{Source: 5, Destination: 10, Length: 5}, result: 15},
		{i: 51, m: day5.Map{Source: 98, Destination: 50, Length: 2}, result: 0, err: true},
	}

	for _, test := range tests {
		result, err := day5.GetRelationForMap(test.i, test.m)
		if err != nil && !test.err {
			t.Errorf("GetRelationForMap(%v, %v) = %v; wanted %v", test.i, test.m, err, test.result)
		}
		if result != test.result {
			t.Errorf("GetRelationForMap(%v, %v) = %v; wanted %v", test.i, test.m, result, test.result)
		}
	}
}

func TestGetRelationForMapSet(t *testing.T) {
	type testArgs struct {
		i      int
		mapSet []day5.Map
		result int
	}

	tests := []testArgs{
		{
			i: 79,
			mapSet: []day5.Map{
				{Destination: 50, Source: 98, Length: 2},
				{Destination: 52, Source: 50, Length: 48},
			},
			result: 81,
		},
		{
			i: 14,
			mapSet: []day5.Map{
				{Destination: 50, Source: 98, Length: 2},
				{Destination: 52, Source: 50, Length: 48},
			},
			result: 14,
		},
		{
			i: 55,
			mapSet: []day5.Map{
				{Destination: 50, Source: 98, Length: 2},
				{Destination: 52, Source: 50, Length: 48},
			},
			result: 57,
		},
		{
			i: 13,
			mapSet: []day5.Map{
				{Destination: 50, Source: 98, Length: 2},
				{Destination: 52, Source: 50, Length: 48},
			},
			result: 13,
		},
	}

	for _, test := range tests {
		result := day5.GetRelationForMapSet(test.i, test.mapSet)
		if result != test.result {
			t.Errorf("GetRelationForMapSet(%v, %v) = %v; wanted %v", test.i, test.mapSet, result, test.result)
		}
	}
}

func TestGetLocationsForSeedsIndividual(t *testing.T) {
	result := day5.GetLocationsForSeedsIndividual("test-input.txt")
	expected := []int{82, 43, 86, 35}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetLocationsForSeeds() = %v; wanted %v", result, expected)
	}
}

func TestGetLocationsForSeedRanges(t *testing.T) {
	result := day5.GetLocationForSeedRanges("test-input.txt")
	expected := 46
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetLocationsForSeeds() = %v; wanted %v", result, expected)
	}
}

func TestGetLocationsForSeedRanges2(t *testing.T) {
	result := day5.GetLocationForSeedRanges2("test-input.txt")
	expected := 46
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetLocationsForSeeds() = %v; wanted %v", result, expected)
	}
}
