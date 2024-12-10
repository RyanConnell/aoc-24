package main

import (
	"fmt"

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

type Point struct{ X, Y int }

func parseInput(lines []string) ([][]int, []Point) {
	world := make([][]int, len(lines))
	var starts []Point
	for y, line := range lines {
		world[y] = make([]int, len(line))
		for x, char := range line {
			world[y][x] = int(char - '0')
			if char == '0' {
				starts = append(starts, Point{X: x, Y: y})
			}
		}
	}
	return world, starts
}

/// Part 1 \\\

func solve(lines []string) (int, error) {
	world, starts := parseInput(lines)
	var sum int
	for _, start := range starts {
		results := make(map[Point]struct{})
		visited := make([][]bool, len(lines))
		for y, line := range lines {
			visited[y] = make([]bool, len(line))
		}
		findUniqueEndpoints(start.X, start.Y, 0, world, visited, results)
		sum += len(results)
	}
	return sum, nil
}

func findUniqueEndpoints(x, y, currentHeight int, world [][]int, visited [][]bool, results map[Point]struct{}) {
	if visited[y][x] {
		return
	}
	visited[y][x] = true
	if world[y][x] == 9 {
		results[Point{X: x, Y: y}] = struct{}{}
		return
	}
	if world[y][x] != currentHeight {
		return
	}
	if canMove(x+1, y, currentHeight+1, world) {
		findUniqueEndpoints(x+1, y, currentHeight+1, world, visited, results)
	}
	if canMove(x-1, y, currentHeight+1, world) {
		findUniqueEndpoints(x-1, y, currentHeight+1, world, visited, results)
	}
	if canMove(x, y+1, currentHeight+1, world) {
		findUniqueEndpoints(x, y+1, currentHeight+1, world, visited, results)
	}
	if canMove(x, y-1, currentHeight+1, world) {
		findUniqueEndpoints(x, y-1, currentHeight+1, world, visited, results)
	}
}

func canMove(x, y, expectedHeight int, world [][]int) bool {
	if x < 0 || y < 0 || y >= len(world) || x >= len(world[y]) {
		return false
	}
	return world[y][x] == expectedHeight
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	world, starts := parseInput(lines)
	var sum int
	for _, start := range starts {
		results := make([][]int, len(lines))
		for y, line := range lines {
			results[y] = make([]int, len(line))
			for x := range lines {
				results[y][x] = -1
			}
		}
		sum += sumPaths(start.X, start.Y, 0, world, results)
	}
	return sum, nil
}

func sumPaths(x, y, currentHeight int, world [][]int, results [][]int) int {
	if currentHeight == 9 {
		return 1
	}
	if results[y][x] != -1 {
		return results[y][x]
	}
	var sum int
	if canMove(x+1, y, currentHeight+1, world) {
		sum += sumPaths(x+1, y, currentHeight+1, world, results)
	}
	if canMove(x-1, y, currentHeight+1, world) {
		sum += sumPaths(x-1, y, currentHeight+1, world, results)
	}
	if canMove(x, y+1, currentHeight+1, world) {
		sum += sumPaths(x, y+1, currentHeight+1, world, results)
	}
	if canMove(x, y-1, currentHeight+1, world) {
		sum += sumPaths(x, y-1, currentHeight+1, world, results)
	}
	results[y][x] = sum
	return sum
}
