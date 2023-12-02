package day1

import (
	"aoc2023/helpers"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var threeNumbers = []string{"one", "two", "six"}
var fourNumbers = append([]string{"four", "five", "nine"}, threeNumbers...)
var fiveNumbers = append([]string{"three", "seven", "eight"}, fourNumbers...)

func parseIntWord(s string) (int, error) {
	for i, num := range numbers {
		if s == num {
			return i, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("%v is not a number", s))
}

func findSubstringInt(arr []string, s string) (int, error) {
	for _, val := range arr {
		if strings.Contains(s, val) {
			return parseIntWord(val)
		}
	}
	return 0, errors.New("no int substring found")
}

func getIntForSlice(slice string) (int, error) {
	switch len(slice) {
	case 1:
		return 0, errors.New(fmt.Sprintf("Can't find int for unsupported length %v", slice))
	case 2:
		return 0, errors.New(fmt.Sprintf("Can't find int for unsupported length %v", slice))
	case 3:
		return findSubstringInt(threeNumbers, slice)
	case 4:
		return findSubstringInt(fourNumbers, slice)
	default:
		return findSubstringInt(fiveNumbers, slice)
	}
}

func findValue(line string, withWords bool) (int, error) {
	first, last := 10, 10

	charSlice := strings.Split(line, "")
	startAcc := []string{}
	endAcc := []string{}
	for _, char := range charSlice {
		if val, err := strconv.Atoi(char); err == nil {
			first = val
			break
		}

		if withWords {
			startAcc = append(startAcc, char)
			if val, err := getIntForSlice(strings.Join(startAcc, "")); err == nil {
				first = val
				break
			}
		}
	}

	for k := len(charSlice) - 1; k >= 0; k = k - 1 {
		char := charSlice[k]
		if val, err := strconv.Atoi(char); err == nil {
			last = val
			break
		}

		if withWords {
			endAcc = append([]string{char}, endAcc...)
			if val, err := getIntForSlice(strings.Join(endAcc, "")); err == nil {
				last = val
				break
			}
		}
	}
	if first < 10 && last < 10 {
		return first*10 + last, nil
	}
	return 0, errors.New(fmt.Sprintf("no value found for %s", line))
}

func findTotalForLines(input string, withWords bool) int {
	lines := helpers.ReadInputToLines(input)

	values := []int{}
	for i := 0; i < len(lines); i = i + 1 {
		line := lines[i]
		val, err := findValue(line, withWords)
		if err == nil {
			values = append(values, val)
		} else {
			fmt.Println(err)
		}
	}

	return helpers.Sum(values)
}

func Part1() string {
	total := findTotalForLines("day1/input.txt", false)
	return fmt.Sprint(total)
}

func Part2() string {
	total := findTotalForLines("day1/input.txt", true)
	return fmt.Sprint(total)
}
