package main

import (
	"fmt"
	"regexp"
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
	expr := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	return sumLines(lines, expr)
}

func sumLines(lines []string, expr *regexp.Regexp) (int, error) {
	var sum int
	do := true
	for _, line := range lines {
		expressions := expr.FindAllStringSubmatch(line, -1)
		for _, ex := range expressions {
			if ex[0] == "don't()" {
				do = false
			}
			if ex[0] == "do()" {
				do = true
			}
			if do && strings.HasPrefix(ex[0], "mul(") {
				left, err := strconv.Atoi(ex[1])
				if err != nil {
					return 0, err
				}
				right, err := strconv.Atoi(ex[2])
				if err != nil {
					return 0, err
				}
				sum += left * right
			}
		}
	}
	return sum, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	expr := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|do\\(\\)|don't\\(\\)")
	return sumLines(lines, expr)
}
