package day6

import (
	"aoc2023/helpers"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func getIntValues(line string) []int {
	re := regexp.MustCompile(` +`)
	clean := re.ReplaceAllString(line, " ")
	nums := strings.Split(clean, " ")[1:]
	return helpers.Map(nums, func(num string, _ int) int {
		if num, err := strconv.Atoi(num); err == nil {
			return num
		}
		fmt.Printf("Failed to parse num %v\n", num)
		return 0
	})
}

type Race struct {
	Time           int
	DistanceRecord int
}

func buildRaces(times []int, distances []int) []Race {
	races := []Race{}
	for i, t := range times {
		races = append(races, Race{Time: t, DistanceRecord: distances[i]})
	}
	return races
}

func GetMinHoldTime(time int, distanceRecord int) int {
	timeF := float64(time)
	distanceRecordF := float64(distanceRecord)
	min := math.Floor(distanceRecordF / timeF)
	for n := min; n*(timeF-n) <= distanceRecordF; n = n + 1 {
		min = n
	}
	return int(min + 1)
}

func GetMaxHoldTime(time int, distanceRecord int, min int) int {
	max := float64(min + 1)
	timeF := float64(time)
	distanceRecordF := float64(distanceRecord)
	for n := max; n*(timeF-n) > distanceRecordF; n = n + 1 {
		max = n
	}
	return int(max)
}

func GetWinStrategyCount(race Race) int {
	min := GetMinHoldTime(race.Time, race.DistanceRecord)
	max := GetMaxHoldTime(race.Time, race.DistanceRecord, min)
	return max - min + 1
}

func readAndParseRaces(input string) []Race {
	lines := helpers.ReadInputToLines(input)
	timePart := lines[0]
	distancePart := lines[1]

	times := getIntValues(timePart)
	distances := getIntValues(distancePart)

	return buildRaces(times, distances)
}

func GetWinStrategyProduct(input string) int {
	races := readAndParseRaces(input)
	return helpers.Reduce(races, func(result int, race Race, _ int) int {
		count := GetWinStrategyCount(race)
		return result * count
	}, 1)
}

func Part1() string {
	result := GetWinStrategyProduct("day6/input.txt")
	return fmt.Sprint(result)
}

func combineRaces(races []Race) Race {
	timeS := helpers.Reduce(races, func(result string, race Race, _ int) string {
		return fmt.Sprintf("%v%v", result, race.Time)
	}, "")
	distanceS := helpers.Reduce(races, func(result string, race Race, _ int) string {
		return fmt.Sprintf("%v%v", result, race.DistanceRecord)
	}, "")
	time, err := strconv.Atoi(timeS)
	if err != nil {
		fmt.Printf("failed to parse combined time %v\n", timeS)
	}
	distance, err := strconv.Atoi(distanceS)
	if err != nil {
		fmt.Printf("failed to parse combined distance %v\n", distanceS)
	}
	return Race{Time: time, DistanceRecord: distance}
}

func GetWinStrategyCountOfCombinedRace(input string) int {
	races := readAndParseRaces(input)
	race := combineRaces(races)
	return GetWinStrategyCount(race)
}

func Part2() string {
	result := GetWinStrategyCountOfCombinedRace("day6/input.txt")
	return fmt.Sprint(result)
}
