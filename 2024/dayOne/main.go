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

	fmt.Println("Day One Part One Solution: ", part1Solution)
}

func fetchInput() string {
	input, err := os.ReadFile("../inputs/dayOne.txt")
	if err != nil {
		panic("unable to read file!")
	}
	return string(input)
}

func part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	left, right := parseAndSortLists(lines)
	return calculateDistance(left, right)
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
