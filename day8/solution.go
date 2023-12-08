package day8

import (
	"aoc2023/helpers"
	"fmt"
	"strings"
)

type Node struct {
	L       string
	R       string
	Name    string
	IsStart bool
	IsEnd   bool
}
type NodeMap map[string]Node

func parseInput(input string) ([]string, NodeMap, []Node) {
	raw := helpers.ReadInput(input)
	parts := strings.Split(raw, "\n\n")

	instructionsPart := parts[0]
	instructions := strings.Split(instructionsPart, "")

	nodesPart := parts[1]
	nodeStrings := strings.Split(nodesPart, "\n")
	nodeMap := NodeMap{}
	nodes := []Node{}

	for _, nodeString := range nodeStrings {
		captureMap := helpers.GetCaptureGroupMap(`(?P<Name>[A-Z1-9]{3}) = \((?P<L>[A-Z1-9]{3}), (?P<R>[A-Z1-9]{3})\)`, nodeString)

		node := Node{
			L:       captureMap["L"],
			R:       captureMap["R"],
			Name:    captureMap["Name"],
			IsStart: strings.HasSuffix(captureMap["Name"], "A"),
			IsEnd:   strings.HasSuffix(captureMap["Name"], "Z"),
		}

		nodes = append(nodes, node)
		nodeMap[node.Name] = node
	}
	return instructions, nodeMap, nodes
}

func countStepsToFinish(current string, nodeMap NodeMap, instructions []string, isEnd func(node Node) bool) int {
	steps := 0
	for true {
		node := nodeMap[current]
		if isEnd(node) {
			break
		}
		instruction := instructions[steps%len(instructions)]
		if instruction == "L" {
			current = node.L
		} else {
			current = node.R
		}
		steps += 1
	}
	return steps
}

func CountStepsToFinish(input string) int {
	instructions, nodeMap, _ := parseInput(input)

	return countStepsToFinish("AAA", nodeMap, instructions, func(node Node) bool { return node.Name == "ZZZ" })
}

func Part1() string {
	return fmt.Sprint(CountStepsToFinish("day8/input.txt"))
}

func CountGhostStepsToFinish(input string) int {
	instructions, nodeMap, nodes := parseInput(input)

	currents := helpers.Filter(nodes, func(node Node) bool {
		return node.IsStart
	})

	stepCounts := helpers.Map(currents, func(node Node, _ int) int {
		return countStepsToFinish(node.Name, nodeMap, instructions, func(node Node) bool { return node.IsEnd })
	})

	return helpers.Reduce(stepCounts, func(result int, stepCount int, i int) int {
		return helpers.LowestCommonMultiple(result, stepCount)
	}, 1)
}

func Part2() string {
	return fmt.Sprint(CountGhostStepsToFinish("day8/input.txt"))
}
