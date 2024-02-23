package day10

import (
	"aoc2023/helpers"
	"strings"
)

var traversableMap = map[string][]string{
	"|":      {"|", "J", "L", "7", "F"},
	"L":      {"|", "J", "7", "F", "-"},
	"J":      {"|", "L", "7", "F", "-"},
	"7":      {"|", "L", "J", "F", "-"},
	"F":      {"|", "L", "J", "7", "-"},
	"-":      {"-", "F", "7", "J", "L"},
	"top":    {"|", "F", "7"},
	"left":   {"-", "L", "F"},
	"bottom": {"|", "L", "J"},
	"right":  {"-", "7", "J"},
}

func walkMap(mapGrid [][]string, start []int) int {
	var a []int
	var b []int
	for x, y := start[0], start[1]; a != nil; {
		top := []int{x, y - 1}

	}
}

func FindFurthestDistance(input string) int {
	lines := helpers.ReadInputToLines(input)
	mapGrid := helpers.Reduce(lines, func(mapGrid [][]string, line string, i int) [][]string {
		return append(mapGrid, strings.Split(line, ""))
	}, [][]string{})

	start := []int{}
	for y, line := range mapGrid {
		for x, char := range line {
			if char == "S" {
				start = []int{x, y}
			}
		}
	}
	furthest := walkMap(mapGrid, start)
}

func Part1() string {
	return ""
}

func Part2() string {
	return ""
}
