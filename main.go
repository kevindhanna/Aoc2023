package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"aoc2023/day4"
	"aoc2023/day5"
	"aoc2023/day6"
	"aoc2023/day7"
	"aoc2023/day8"
	"aoc2023/day9"
	"aoc2023/helpers"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Day func() string

var implementedDays = [][]Day{
	{day1.Part1, day1.Part2},
	{day2.Part1, day2.Part2},
	{day3.Part1, day3.Part2},
	{day4.Part1, day4.Part2},
	{day5.Part1, day5.Part2},
	{day6.Part1, day6.Part2},
	{day7.Part1, day7.Part2},
	{day8.Part1, day8.Part2},
	{day9.Part1, day9.Part2},
}

func solutionContent(day string) string {
	return fmt.Sprintf("package %v\n\nfunc Part1() string {\n return \"\" \n}\n\nfunc Part2() string {\n return \"\" \n}\n\n", day)
}

func buildScaffolding(day string) {
	dirName := fmt.Sprintf("day%v", day)
	os.Mkdir(dirName, 0755)
	files := []struct {
		name    string
		content []byte
	}{
		{"solution.go", []byte(solutionContent(dirName))},
		{"solution_test.go", []byte{}},
		{"input.txt", []byte{}},
		{"test-input.txt", []byte{}},
	}
	for _, file := range files {
		os.WriteFile(fmt.Sprintf("%v/%v", dirName, file.name), file.content, 0755)
	}
}

func main() {
	days := os.Args[1:]
	scafFlag := flag.Bool("scaf", false, "When present, creates the scaffolding for the days listed")

	flag.Parse()

	if *scafFlag {
		for _, dayString := range days {
			buildScaffolding(dayString)
		}
	} else if len(days) != 0 {
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
