package day8_test

import (
	"aoc2023/day8"
	"testing"
)

func TestCountStepsToFinish(t *testing.T) {
	result := day8.CountStepsToFinish("test-input1.txt")
	if result != 2 {
		t.Errorf("CountStepsToFinish() = %v; wanted %v", result, 2)
	}

	result = day8.CountStepsToFinish("test-input2.txt")
	if result != 6 {
		t.Errorf("CountStepsToFinish() = %v; wanted %v", result, 6)
	}
}

func TestCountGhostStepsToFinish(t *testing.T) {
	result := day8.CountGhostStepsToFinish("test-input3.txt")
	if result != 6 {
		t.Errorf("CountStepsToFinish() = %v; wanted %v", result, 6)
	}
}
