package main

import (
	"fmt"
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

func parseRow(line string) ([]int, error) {
	var ints []int
	for _, val := range strings.Split(line, " ") {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func solve(lines []string) (int, error) {
	var safe int
	for _, line := range lines {
		ints, err := parseRow(line)
		if err != nil {
			return 0, err
		}
		if len(ints) == 1 {
			safe++
			continue
		}

		if check(ints, -1) {
			safe++
			continue
		}
	}
	return safe, nil
}

func check(original []int, skip int) bool {
	var ints []int
	if skip == -1 {
		ints = original
	} else {
		ints = append(ints, original[:skip]...)
		if skip+1 < len(original) {
			ints = append(ints, original[skip+1:]...)
		}
	}

	isAsc := ints[1]-ints[0] > 0
	for i := 0; i < len(ints)-1; i++ {
		diff := ints[i+1] - ints[i]

		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if isAsc && diff < 0 {
			return false
		}
		if !isAsc && diff > 0 {
			return false
		}
	}
	return true
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	var safe int
	for _, line := range lines {
		ints, err := parseRow(line)
		if err != nil {
			return 0, err
		}
		if len(ints) == 1 {
			safe++
			continue
		}

		if check(ints, -1) {
			safe++
			continue
		}
		var isSafe bool
		for i := 0; i <= len(ints); i++ {
			if check(ints, i) {
				isSafe = true
				break
			}
		}
		if isSafe {
			safe++
		}
	}
	return safe, nil
}
