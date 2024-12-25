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
	Name      string
	Keys      []string
	Operation string
}

func NewEquation(line string) *Equation {
	parts := strings.Split(line, " ")
	eq := &Equation{
		Name:      parts[4],
		Keys:      []string{parts[0], parts[2]},
		Operation: parts[1],
	}
	return eq
}

func (e *Equation) Possible(cache map[string]bool) bool {
	for _, key := range e.Keys {
		if _, ok := cache[key]; !ok {
			return false
		}
	}
	return true
}

func (e *Equation) Result(cache map[string]bool) bool {
	left := cache[e.Keys[0]]
	right := cache[e.Keys[1]]

	switch e.Operation {
	case "AND":
		return left && right
	case "XOR":
		return left != right
	case "OR":
		return left || right
	}
	panic(fmt.Sprintf("Unknown operation: %v", e.Operation))
}

func solve(lines []string) (int, error) {
	cache := make(map[string]bool)
	var equations []*Equation
	var numZs int
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "->") {
			eq := NewEquation(line)
			equations = append(equations, eq)
			if eq.Name[0] == 'z' {
				numZs++
			}
		} else if line != "" {
			parts := strings.Split(line, ": ")
			cache[parts[0]] = parts[1] == "1"
			if parts[0][0] == 'z' {
				numZs++
			}
		}
	}

	var workDone bool
	for {
		workDone = false
		for _, eq := range equations {
			if _, ok := cache[eq.Name]; !ok && eq.Possible(cache) {
				cache[eq.Name] = eq.Result(cache)
				workDone = true
			}
		}
		if !workDone {
			break
		}
	}

	zs := make([]bool, numZs)
	for key, value := range cache {
		if key[0] == 'z' {
			idx, err := strconv.Atoi(key[1:])
			if err != nil {
				return 0, err
			}
			zs[numZs-1-idx] = value
		}
	}

	var result int
	for _, val := range zs {
		result = result << 1
		if val {
			result += 1
		}
	}

	return result, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
