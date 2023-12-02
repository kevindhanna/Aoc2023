package day1_test

import (
	"fmt"
	"testing"
)

type testArgs struct {
	input     string
	withWords bool
	result    int
}

func TestFindValue(t *testing.T) {
	tests := []testArgs{
		{input: "6l", withWords: false, result: 66},
		{input: "fone6", withWords: true, result: 16},
		{input: "two1nine", withWords: true, result: 29},
	}

	unit := func(test testArgs) error {
		var result, err = findValue(test.input, test.withWords)
		if err != nil {
			t.Error(fmt.Sprintf("findValue(%v, %v) = %v; got error %v", test.input, test.withWords, test.result, err))
		} else if result != test.result {
			t.Error(fmt.Sprintf("findValue(%v, %v) = %v; got %v", test.input, test.withWords, test.result, result))
		}
		return nil
	}

	for _, t := range tests {
		unit(t)
	}
}

func TestFindTotalForLines(t *testing.T) {
	tests := []testArgs{
		{input: "test-input.txt", withWords: false, result: 142},
		{input: "test-input-p2.txt", withWords: true, result: 281},
	}

	unit := func(test testArgs) error {
		var result = findTotalForLines(test.input, test.withWords)
		if result != test.result {
			t.Error(fmt.Sprintf("findValueForLines(%v) = %v; got %v", test.input, test.result, result))
		}
		return nil
	}

	for _, t := range tests {
		unit(t)
	}
}
