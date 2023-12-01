package main

import (
	"aoc2023/day1"
	"aoc2023/helpers"
	"fmt"
	"os"
)

func main() {
	err := day1.Test()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	helpers.RunPart(day1.Part1, "1-1")
	helpers.RunPart(day1.Part2, "1-2")

	os.Exit(0)
}
