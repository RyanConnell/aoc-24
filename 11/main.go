package main

import (
	"fmt"
	"log"
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

var cache = make(map[int]map[int]int)

func rockSum(value, depth int) int {
	if _, ok := cache[depth]; ok {
		if v, ok := cache[depth][value]; ok && v != 0 {
			return v
		}
	} else {
		cache[depth] = make(map[int]int)
	}

	if depth == 0 {
		cache[depth][value] = 1
		return 1
	}

	if value == 0 {
		result := rockSum(1, depth-1)
		cache[depth][value] = result
		return result
	}

	if numDigits(value)%2 == 0 {
		left, right, err := splitInt(value)
		if err != nil {
			log.Fatalf("Uh oh, error: %v", err)
		}
		result := rockSum(left, depth-1) + rockSum(right, depth-1)
		cache[depth][value] = result
		return result
	}

	result := rockSum(value*2024, depth-1)
	cache[depth][value] = result
	return result
}

func splitInt(in int) (int, int, error) {
	s := strconv.Itoa(in)
	left, err := strconv.Atoi(s[:len(s)/2])
	if err != nil {
		return 0, 0, err
	}
	right, err := strconv.Atoi(s[len(s)/2:])
	if err != nil {
		return 0, 0, err
	}
	return left, right, nil
}

func numDigits(num int) int {
	var count int
	for i := num; i != 0; i /= 10 {
		count++
	}
	return count
}

func solve(lines []string) (int, error) {
	if len(lines) != 1 {
		return 0, fmt.Errorf("bad input")
	}

	stoneStrs := strings.Split(lines[0], " ")
	var values []int
	for _, stoneStr := range stoneStrs {
		val, err := strconv.Atoi(stoneStr)
		if err != nil {
			return 0, err
		}
		values = append(values, val)
	}

	var sum int
	for _, value := range values {
		sum += rockSum(value, 25)
	}

	return sum, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	if len(lines) != 1 {
		return 0, fmt.Errorf("bad input")
	}

	stoneStrs := strings.Split(lines[0], " ")
	var values []int
	for _, stoneStr := range stoneStrs {
		val, err := strconv.Atoi(stoneStr)
		if err != nil {
			return 0, err
		}
		values = append(values, val)
	}

	var sum int
	for _, value := range values {
		sum += rockSum(value, 75)
	}

	return sum, nil
}
