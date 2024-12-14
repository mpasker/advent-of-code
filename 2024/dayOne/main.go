package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dayOneInput := fetchInput()
	part1Solution := part1(dayOneInput)
	part2Solution := part2(dayOneInput)

	fmt.Println("Day One Part One Solution: ", part1Solution)
	fmt.Println("Day One Part Two Solution: ", part2Solution)
}

func fetchInput() string {
	input, err := os.ReadFile("../inputs/dayOne.txt")
	if err != nil {
		panic("unable to read file!")
	}
	return string(input)
}

func part1(input string) int {
	lines := formatLines(input)
	left, right := parseAndSortLists(lines)
	return calculateDistance(left, right)
}

func part2(input string) int {
	lines := formatLines(input)
	left, right := parseAndSortLists(lines)

	return calculateSimilarityScore(left, right)
}

func formatLines(input string) []string {
	formattedString := strings.Split(strings.TrimSpace(input), "\n")
	return formattedString
}

func parseAndSortLists(lines []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)
		leftVal, _ := strconv.Atoi(parts[0])
		rightVal, _ := strconv.Atoi(parts[1])

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func calculateDistance(left []int, right []int) int {
	totalDistance := 0

	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			totalDistance += left[i] - right[i]
		} else {
			totalDistance += right[i] - left[i]
		}
	}

	return totalDistance
}

func calculateSimilarityScore(left []int, right []int) int {
	rightMap := make(map[int]int)
	similarityScore := 0
	for _, num := range right {
		rightMap[num]++
	}

	for _, num := range left {
		if value, exists := rightMap[num]; exists {
			similarityScore += num * value
		}
	}
	return similarityScore
}
