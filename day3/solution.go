package day3

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
)

type Coord struct {
	X int
	Y int
}

func GetSurroundingCoords(coord Coord, maxX int, maxY int) []Coord {
	x, y := coord.X, coord.Y
	result := []Coord{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
	return helpers.Filter(result, func(coord Coord) bool {
		return coord.X >= 0 && coord.Y >= 0 && coord.X <= maxX && coord.Y <= maxY
	})
}

func GetRestOfNumber(point int, line []string) (string, []int) {
	numString := line[point]
	foundPoints := []int{point}
	for i := 1; point-i >= 0; i = i + 1 {
		val := line[point-i]
		if _, err := strconv.Atoi(val); err != nil {
			break
		}
		foundPoints = append(foundPoints, point-i)
		numString = val + numString
	}

	for i := 1; point+i < len(line); i = i + 1 {
		val := line[point+i]
		if _, err := strconv.Atoi(val); err != nil {
			break
		}
		foundPoints = append(foundPoints, point+i)
		numString = numString + val
	}

	return numString, foundPoints
}

func parseSchematic(schematic string) ([][]string, []Coord) {
	symbolCoords := []Coord{}
	grid, err := helpers.SplitMap(schematic, "\n", func(line string, y int) ([]string, error) {
		return helpers.SplitMap(line, "", func(char string, x int) (string, error) {
			if _, err := strconv.Atoi(char); err != nil && char != "." {
				symbolCoords = append(symbolCoords, Coord{x, y})
			}
			return char, nil
		})
	})
	if err != nil {
		fmt.Printf("Couldn't build grid, %v\n", err)
	}
	return grid, symbolCoords
}

func SumPartNumbers(input string) int {
	schematic := helpers.ReadInput(input)

	grid, symbolCoords := parseSchematic(schematic)
	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1

	partNumbers := []int{}
	foundCoords := helpers.MakeGrid[bool](len(grid[0]), len(grid))
	for _, coord := range symbolCoords {
		surroundingCoords := GetSurroundingCoords(coord, maxX, maxY)
		for _, surround := range surroundingCoords {
			if foundCoords[surround.Y][surround.X] {
				continue
			}

			_, err := strconv.Atoi(grid[surround.Y][surround.X])
			if err == nil {
				numString, foundPoints := GetRestOfNumber(surround.X, grid[surround.Y])
				number, err := strconv.Atoi(numString)
				if err != nil {
					fmt.Printf("failed to parse numString %v at %v", numString, surround)
				}
				partNumbers = append(partNumbers, number)
				for _, point := range foundPoints {
					foundCoords[surround.Y][point] = true
				}
			}
		}
	}

	return helpers.Sum(partNumbers)
}

func SumGearRatios(input string) int {
	schematic := helpers.ReadInput(input)

	grid, symbolCoords := parseSchematic(schematic)
	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1

	gearCoords := helpers.Filter(symbolCoords, func(coord Coord) bool {
		return grid[coord.Y][coord.X] == "*"
	})

	gearRatios := []int{}
	for _, coord := range gearCoords {
		foundCoords := helpers.MakeGrid[bool](len(grid[0]), len(grid))
		surroundingCoords := GetSurroundingCoords(coord, maxX, maxY)

		gearNumbers := []int{}
		for _, surround := range surroundingCoords {
			if foundCoords[surround.Y][surround.X] {
				continue
			}
			_, err := strconv.Atoi(grid[surround.Y][surround.X])
			if err == nil {
				numString, foundPoints := GetRestOfNumber(surround.X, grid[surround.Y])
				number, err := strconv.Atoi(numString)
				if err != nil {
					fmt.Printf("failed to parse numString %v at %v", numString, surround)
				}

				gearNumbers = append(gearNumbers, number)
				for _, point := range foundPoints {
					foundCoords[surround.Y][point] = true
				}
			}
		}

		if len(gearNumbers) == 2 {
			gearRatios = append(gearRatios, gearNumbers[0]*gearNumbers[1])
		}
	}

	return helpers.Sum(gearRatios)
}

func Part1() string {
	return fmt.Sprintf("%v", SumPartNumbers("day3/input.txt"))
}

func Part2() string {
	return fmt.Sprintf("%v", SumGearRatios("day3/input.txt"))
}
