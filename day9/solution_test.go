package day9_test

import (
	"aoc2023/day9"
	"testing"
)

func TestSumExtrapolate(t *testing.T) {
	result := day9.SumExtrapolate("test-input.txt")
	if result != 114 {
		t.Errorf("SumExtrapolate() = %v; wanted %v", result, 114)
	}
}

func TestSumExtrapolateBackwards(t *testing.T) {
	result := day9.SumExtrapolateBackwards("test-input.txt")
	if result != 2 {
		t.Errorf("SumExtrapolate() = %v; wanted %v", result, 2)
	}
}
