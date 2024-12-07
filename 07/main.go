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
	equation := Equation{result: result}
	for _, part := range parts[1:] {
		field, err := strconv.Atoi(part)
		if err != nil {
			return Equation{}, err
		}
		equation.fields = append(equation.fields, field)
	}
	return equation, nil
}

func squashEquation(expected int, parts []int, allowConcat bool) bool {
	if len(parts) == 0 {
		return expected == 0
	}
	if len(parts) == 1 {
		return expected == parts[0]
	}

	// Addition
	if squashEquation(expected, append([]int{parts[0] + parts[1]}, parts[2:]...), allowConcat) {
		return true
	}
	if squashEquation(expected, append([]int{parts[0] * parts[1]}, parts[2:]...), allowConcat) {
		return true
	}
	if allowConcat {
		concat, err := strconv.Atoi(fmt.Sprintf("%d%d", parts[0], parts[1]))
		if err != nil {
			panic(err)
		}
		if squashEquation(expected, append([]int{concat}, parts[2:]...), allowConcat) {
			return true
		}
	}
	return false
}

func solve(lines []string) (int, error) {
	var equations []Equation
	for _, line := range lines {
		eq, err := NewEquation(line)
		if err != nil {
			return 0, err
		}
		equations = append(equations, eq)
	}

	var sum int
	for _, eq := range equations {
		solveable := squashEquation(eq.result, eq.fields, false)
		if solveable {
			sum += eq.result
		}
	}
	return sum, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	var equations []Equation
	for _, line := range lines {
		eq, err := NewEquation(line)
		if err != nil {
			return 0, err
		}
		equations = append(equations, eq)
	}

	var sum int
	for _, eq := range equations {
		solveable := squashEquation(eq.result, eq.fields, true)
		if solveable {
			sum += eq.result
		}
	}
	return sum, nil
}
