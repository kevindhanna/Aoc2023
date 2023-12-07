package day5

import (
	"aoc2023/helpers"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseSeeds(seedsPart string) []int {
	seedStrings := strings.Split(seedsPart, " ")[1:]
	return helpers.Map(seedStrings, func(seed string, _ int) int {
		if val, err := strconv.Atoi(seed); err == nil {
			return val
		}
		fmt.Printf("Failed to parse seed %v\n", seed)
		return 0
	})
}

type Map struct {
	Destination int
	Length      int
	Source      int
}

func ParseMapLine(line string, reverse bool) Map {
	nums, err := helpers.SplitMap(line, " ", func(num string, _ int) (int, error) {
		if val, err := strconv.Atoi(num); err == nil {
			return val, nil
		}
		return 0, errors.New(fmt.Sprintf("Failed to parse map num %v", num))
	})
	if err != nil {
		fmt.Println(errors.New(fmt.Sprintf("Failed to parse map line %v\n", line)))
		return Map{Destination: 0, Source: 0, Length: 0}
	}
	if reverse {
		return Map{
			Destination: nums[1],
			Source:      nums[0],
			Length:      nums[2],
		}
	}
	return Map{
		Destination: nums[0],
		Source:      nums[1],
		Length:      nums[2],
	}
}

func parseMap(mapPart string, reverse bool) []Map {
	partStrings := strings.Split(mapPart, "\n")[1:]
	parts := helpers.Map(partStrings, func(line string, _ int) Map {
		return ParseMapLine(line, reverse)
	})

	return parts
}

func GetRelationForMap(i int, m Map) (int, error) {
	if numInRange(i, m.Source, m.Length) {
		return m.Destination + (i - m.Source), nil
	}
	return 0, errors.New(fmt.Sprintf("No match for %v in map %v", i, m))
}

func GetRelationForMapSet(i int, mapSet []Map) int {
	for _, m := range mapSet {
		v, err := GetRelationForMap(i, m)
		if err == nil {
			return v
		}
	}

	return i
}

func getLocationForSeed(i int, maps [][]Map) int {
	return helpers.Reduce(maps, func(next int, m []Map, _ int) int {
		result := GetRelationForMapSet(next, m)
		return result
	}, i)
}

func parseInput(input string, reverse bool) ([][]Map, []int) {
	text := helpers.ReadInput(input)

	parts := strings.Split(text, "\n\n")
	seedsPart := parts[0]
	seedToSoilPart := parts[1]
	soilToFertilizerPart := parts[2]
	fertilizerToWaterPart := parts[3]
	waterToLightPart := parts[4]
	lightToTemperaturePart := parts[5]
	temperatureToHumidityPart := parts[6]
	humidityToLocationPart := parts[7]

	seeds := parseSeeds(seedsPart)
	seedToSoilMap := parseMap(seedToSoilPart, reverse)
	soilToFertilizerMap := parseMap(soilToFertilizerPart, reverse)
	fertilizerToWaterMap := parseMap(fertilizerToWaterPart, reverse)
	waterToLightMap := parseMap(waterToLightPart, reverse)
	lightToTemperatureMap := parseMap(lightToTemperaturePart, reverse)
	temperatureToHumidityMap := parseMap(temperatureToHumidityPart, reverse)
	humidityToLocationMap := parseMap(humidityToLocationPart, reverse)

	return [][]Map{
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	}, seeds
}

func GetLocationsforSeeds(seeds []int, maps [][]Map) []int {
	locations := helpers.Map(seeds, func(seed int, _ int) int {
		return getLocationForSeed(seed, maps)
	})

	return locations
}

func GetLocationsForSeedsIndividual(input string) []int {
	maps, seeds := parseInput(input, false)
	locations := GetLocationsforSeeds(seeds, maps)
	return locations
}

func Part1() string {
	locations := GetLocationsForSeedsIndividual("day5/input.txt")
	return fmt.Sprint(helpers.Min(locations))
}

func GetLocationForSeedRanges(input string) int {
	maps, seedRanges := parseInput(input, false)

	min := -1
	r := 1
	for i := 0; i < len(seedRanges); i = i + 2 {
		start := seedRanges[i]
		length := seedRanges[i+1]
		end := start + length
		fmt.Printf("r = %v\n", r)

		for i := start; i <= end; i = i + 1 {
			loc := getLocationForSeed(i, maps)
			if min == -1 || min > loc {
				min = loc
				fmt.Printf("min = %v\n", min)
			}
		}
		r = r + 1
	}
	return min
}

func numInRange(i int, start int, length int) bool {
	return i >= start && i <= start+length
}

func GetLocationForSeedRanges2(input string) int {
	maps, seedRanges := parseInput(input, true)
	maps = helpers.Reverse(maps)
	locationMap := maps[0]
	lowestLocationMap := locationMap[0]
	for i := 1; i < len(maps[0]); i = i + 1 {
		current := locationMap[i]
		if current.Source < lowestLocationMap.Source {
			lowestLocationMap = current
		}
	}

	for i := lowestLocationMap.Source; i < lowestLocationMap.Source+lowestLocationMap.Length; i = i + 1 {
		fmt.Printf("i = %v\n", i)
		loc := getLocationForSeed(i, maps)
		fmt.Printf("")
		for seedI := 0; seedI < len(seedRanges); seedI = seedI + 2 {
			if numInRange(loc, seedRanges[seedI], seedRanges[seedI+1]) {
				return i
			}
		}
	}

	return -1
}

func Part2() string {
	min := GetLocationForSeedRanges2("day5/input.txt")
	return fmt.Sprint(min)
}
