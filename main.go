package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"aoc2023/day4"
	"aoc2023/day5"
	"aoc2023/day6"
	"aoc2023/helpers"
	"fmt"
	"os"
	"strconv"
)

type Day func() string

var implementedDays = [][]Day{{day1.Part1, day1.Part2},
	{day2.Part1, day2.Part2},
	{day3.Part1, day3.Part2},
	{day4.Part1, day4.Part2},
	{day5.Part1, day5.Part2},
	{day6.Part1, day6.Part2},
}

func main() {
	days := os.Args[1:]

	if len(days) != 0 {
		for _, dayString := range days {
			day, err := strconv.Atoi(dayString)
			if err != nil {
				fmt.Println(fmt.Errorf("%v is not a day, man", dayString))
				return
			}
			if day > len(implementedDays) {
				fmt.Println(fmt.Errorf("Day %v not implemented yet, dude", day))
				return
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
