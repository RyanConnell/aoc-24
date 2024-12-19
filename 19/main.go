package main

import (
	"fmt"
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
	blocks := strings.Split(lines[0], ", ")
	var count int
	for _, line := range lines[2:] {
		cache := make(map[string]int)
		if possible(line, blocks, cache) != 0 {
			count++
		}
	}
	return count, nil
}

func possible(word string, blocks []string, cache map[string]int) int {
	if result, ok := cache[word]; ok {
		return result
	}

	var uniqueBlocks int
	for _, block := range blocks {
		if len(block) > len(word) {
			continue
		}
		if block == word {
			uniqueBlocks += 1
			continue
		}
		if block == word[:len(block)] {
			result := possible(word[len(block):], blocks, cache)
			uniqueBlocks += result
		}
	}
	cache[word] = uniqueBlocks
	return uniqueBlocks
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	blocks := strings.Split(lines[0], ", ")
	var sum int
	for _, line := range lines[2:] {
		cache := make(map[string]int)
		sum += possible(line, blocks, cache)
	}
	return sum, nil
}
