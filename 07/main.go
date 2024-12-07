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

type Equation struct {
	fields []int
	result int
}

func NewEquation(line string) (Equation, error) {
	parts := strings.Split(line, " ")
	result, err := strconv.Atoi(parts[0][:len(parts[0])-1])
	if err != nil {
		return Equation{}, err
	}
	equation := Equation{result: result, fields: make([]int, len(parts)-1)}
	for i, part := range parts[1:] {
		field, err := strconv.Atoi(part)
		if err != nil {
			return Equation{}, err
		}
		equation.fields[i] = field
	}
	return equation, nil
}

func squashEquation(expected int, firstElem int, parts []int, allowConcat bool) bool {
	if len(parts) == 0 {
		return expected == firstElem
	}
	if expected < firstElem {
		return false
	}

	if squashEquation(expected, firstElem+parts[0], parts[1:], allowConcat) {
		return true
	}
	if squashEquation(expected, firstElem*parts[0], parts[1:], allowConcat) {
		return true
	}
	if allowConcat {
		concat := concatNums(firstElem, parts[0])
		if squashEquation(expected, concat, parts[1:], allowConcat) {
			return true
		}
	}
	return false
}

func concatNums(left, right int) int {
	for i := right; i != 0; i /= 10 {
		left *= 10
	}
	return left + right
}

func solve(lines []string) (int, error) {
	return solution(lines, false)
}

func solution(lines []string, allowConcat bool) (int, error) {
	var sum int
	for idx := range lines {
		eq, _ := NewEquation(lines[idx])
		solveable := squashEquation(eq.result, eq.fields[0], eq.fields[1:], allowConcat)
		if solveable {
			sum += eq.result
		}
	}
	return sum, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return solution(lines, true)
}
