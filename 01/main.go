package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

func solve(lines []string) (int, error) {
	var diff int
	var lefts, rights []int
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		lefts = append(lefts, left)
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		rights = append(rights, right)
	}
	sort.Ints(lefts)
	sort.Ints(rights)

	for i := 0; i < len(lefts); i++ {
		diff += int(math.Abs(float64(lefts[i] - rights[i])))
	}
	return diff, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	var lefts []int
	rightCounts := make(map[int]int)
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		lefts = append(lefts, left)
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		rightCounts[right]++
	}

	var diff int
	for _, left := range lefts {
		diff += left * rightCounts[left]
	}
	return diff, nil
}
