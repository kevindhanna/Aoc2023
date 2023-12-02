package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/helpers"
	"fmt"
	"os"
	"strconv"
)

type Day func() string

var implementedDays = [][]Day{{day1.Part1, day1.Part2}, {day2.Part1, day2.Part2}}

func main() {
	days := os.Args[1:]

	if len(days) != 0 {
		for _, dayString := range days {
			day, err := strconv.Atoi(dayString)
			if err != nil {
				fmt.Errorf("%v is not a day, man", dayString)
			}
			if day > len(days) {
				fmt.Errorf("Day %v not implemented yet, dude", day)
			}
			dayFns := implementedDays[day-1]
			for i, fn := range dayFns {
				helpers.RunPart(fn, fmt.Sprintf("%v:%v", day, i+1))
			}
		}
	} else {
		for i, fns := range implementedDays {
			for j, fn := range fns {
				helpers.RunPart(fn, fmt.Sprintf("%v:%v", i+1, j+1))
			}
		}
	}

	os.Exit(0)
}
