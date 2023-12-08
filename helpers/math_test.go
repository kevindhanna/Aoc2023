package helpers_test

import (
	"aoc2023/helpers"
	"testing"
)

func TestLowestCommonMultiple(t *testing.T) {
	result := helpers.LowestCommonMultiple(33, 75)
	if result != 825 {
		t.Errorf("LowestCommonMultiple(33, 75) = %v; wanted %v", result, 825)
	}
}

func TestGreatestCommonFactor(t *testing.T) {
	result := helpers.GreatestCommonFactor(50, 15)
	if result != 5 {
		t.Errorf("LowestCommonMultiple(50, 15) = %v; wanted %v", result, 5)
	}
}
