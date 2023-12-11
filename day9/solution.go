package day9

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
	"strings"
)

func Extrapolate(historySet [][]int) [][]int {
	diffs := []int{}
	history := historySet[len(historySet)-1]
	for i := 0; len(diffs) != len(history)-1; i += 1 {
		a, b := history[i], history[i+1]
		diffs = append(diffs, b-a)
	}
	historySet = append(historySet, diffs)
	if helpers.Every(historySet[len(historySet)-1], func(val int) bool {
		return val == 0
	}) {
		return historySet
	}
	return Extrapolate(historySet)
}

func FindNext(history []int) int {
	extrapolated := Extrapolate([][]int{history})

	// populate the new zero on the last set
	extrapolated[len(extrapolated)-1] = append(extrapolated[len(extrapolated)-1], 0)
	for i := len(extrapolated) - 2; i >= 0; i -= 1 {
		current := extrapolated[i]
		previous := extrapolated[i+1]
		diff := previous[len(previous)-1]
		extrapolated[i] = append(current, current[len(current)-1]+diff)
	}

	return extrapolated[0][len(extrapolated[0])-1]
}

func SumExtrapolate(input string) int {
	historyStrings := helpers.ReadInputToLines(input)
	histories := helpers.Map(historyStrings, func(historyString string, i int) []int {
		nums := strings.Split(historyString, " ")
		return helpers.Map(nums, func(num string, i int) int {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Failed to parse num %v", num)
			}
			return val
		})
	})
	return helpers.Reduce(histories, func(total int, history []int, i int) int {
		return total + FindNext(history)
	}, 0)
}

func SumExtrapolateBackwards(input string) int {
	historyStrings := helpers.ReadInputToLines(input)
	histories := helpers.Map(historyStrings, func(historyString string, i int) []int {
		nums := strings.Split(historyString, " ")
		history := helpers.Map(nums, func(num string, i int) int {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Failed to parse num %v", num)
			}
			return val
		})
		return helpers.Reverse(history)
	})
	return helpers.Reduce(histories, func(total int, history []int, i int) int {
		return total + FindNext(history)
	}, 0)
}

func Part1() string {
	return fmt.Sprint(SumExtrapolate(("day9/input.txt")))
}

func Part2() string {
	return fmt.Sprint(SumExtrapolateBackwards(("day9/input.txt")))
}
