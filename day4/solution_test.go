package day4_test

import (
	"aoc2023/day4"
	"testing"
)

func TestSumWinningCardPoints(t *testing.T) {
	result := day4.SumWinningCardPoints("test-input.txt")
	expected := 13
	if result != 13 {
		t.Errorf("SumWinningCardPoints() = %v, wanted %v", result, expected)
	}
}

func TestSumTotalCards(t *testing.T) {
	result := day4.SumTotalCards("test-input.txt")
	expected := 30
	if result != expected {
		t.Errorf("SumTotalCards() = %v; wanted %v", result, expected)
	}
}
