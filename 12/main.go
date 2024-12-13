package main

import (
	"fmt"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/sample.txt")

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
	results := make(map[int]*Result)
	regions, visited := parseInput(lines)

	// Count the area and borders of each region.
	for y, line := range regions {
		for x := range line {
			quantifyRegion(x, y, regions, visited, results)
		}
	}

	var sum int
	for _, result := range results {
		sum += result.Area * result.Border
	}

	return sum, nil
}

func floodLand(land [][]rune) [][]int {
	regions := make([][]int, len(land))
	for y := 0; y < len(regions); y++ {
		regions[y] = make([]int, len(land[y]))
	}
	var flood func(id, x, y int)

	flood = func(id, x, y int) {
		if regions[y][x] != 0 {
			return
		}
		regions[y][x] = id
		if inBounds(x, y-1, regions) && land[y][x] == land[y-1][x] {
			flood(id, x, y-1)
		}
		if inBounds(x, y+1, regions) && land[y][x] == land[y+1][x] {
			flood(id, x, y+1)
		}
		if inBounds(x-1, y, regions) && land[y][x] == land[y][x-1] {
			flood(id, x-1, y)
		}
		if inBounds(x+1, y, regions) && land[y][x] == land[y][x+1] {
			flood(id, x+1, y)
		}
	}

	id := 1
	for y, line := range regions {
		for x := range line {
			if regions[y][x] != 0 {
				continue
			}
			id++
			flood(id, x, y)
		}
	}

	return regions
}

func zoom(regions [][]int) [][]int {
	zoomedRegions := make([][]int, len(regions)*3)
	for y, line := range regions {
		zoomedRegions[(y*3)+0] = make([]int, len(line)*3)
		zoomedRegions[(y*3)+1] = make([]int, len(line)*3)
		zoomedRegions[(y*3)+2] = make([]int, len(line)*3)
		for x, id := range line {
			if !inBounds(x, y-1, regions) || regions[y][x] != regions[y-1][x] {
				// Draw line at top of tile.
				zoomedRegions[(y*3)+0][(x*3)+0] = id
				zoomedRegions[(y*3)+0][(x*3)+1] = id
				zoomedRegions[(y*3)+0][(x*3)+2] = id
			}
			if !inBounds(x, y+1, regions) || regions[y][x] != regions[y+1][x] {
				// Draw line at bottom of tile.
				zoomedRegions[(y*3)+2][(x*3)+0] = id
				zoomedRegions[(y*3)+2][(x*3)+1] = id
				zoomedRegions[(y*3)+2][(x*3)+2] = id
			}
			if !inBounds(x-1, y, regions) || regions[y][x] != regions[y][x-1] {
				// Draw line at left of tile.
				zoomedRegions[(y*3)+0][(x*3)+0] = id
				zoomedRegions[(y*3)+1][(x*3)+0] = id
				zoomedRegions[(y*3)+2][(x*3)+0] = id
			}
			if !inBounds(x+1, y, regions) || regions[y][x] != regions[y][x+1] {
				// Draw line at left of tile.
				zoomedRegions[(y*3)+0][(x*3)+2] = id
				zoomedRegions[(y*3)+1][(x*3)+2] = id
				zoomedRegions[(y*3)+2][(x*3)+2] = id
			}
		}
	}
	return zoomedRegions
}

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

func countEdges(x, y int, zoomed [][]int) int {
	target := zoomed[y][x]
	if target == 0 {
		return 0
	}

	var edgeCount, steps int
	var xMove, yMove int
	direction := NORTH
	for inBounds(y, x, zoomed) && zoomed[y][x] == target {
		switch direction {
		case NORTH:
			xMove, yMove = 0, -1
		case SOUTH:
			xMove, yMove = 0, 1
		case EAST:
			xMove, yMove = 1, 0
		case WEST:
			xMove, yMove = -1, 0
		}

		if !inBounds(x+xMove, y+yMove, zoomed) || zoomed[y+yMove][x+xMove] != target {
			if steps != 0 {
				edgeCount++
				steps = 0
			}
			// Choose next direction
			if inBounds(x, y-1, zoomed) && zoomed[y-1][x] == target {
				direction = NORTH
			} else if inBounds(x+1, y, zoomed) && zoomed[y][x+1] == target {
				direction = EAST
			} else if inBounds(x-1, y, zoomed) && zoomed[y][x-1] == target {
				direction = WEST
			} else if inBounds(x, y+1, zoomed) && zoomed[y+1][x] == target {
				direction = SOUTH
			} else {
				break
			}
			continue
		}

		zoomed[y][x] = 0
		steps++
		y += yMove
		x += xMove
	}
	return edgeCount
}

type Result struct {
	Area   int
	Border int
}

func quantifyRegion(x, y int, land [][]int, visited [][]bool, results map[int]*Result) {
	if !inBounds(x, y, land) || visited[y][x] {
		return
	}

	self := land[y][x]
	if _, ok := results[self]; !ok {
		results[self] = &Result{Area: 1}
	} else {
		results[self].Area++
	}

	// Check above
	if !inBounds(x, y-1, land) || land[y][x] != land[y-1][x] {
		results[self].Border++
	}
	if !inBounds(x, y+1, land) || land[y][x] != land[y+1][x] {
		results[self].Border++
	}
	if !inBounds(x-1, y, land) || land[y][x] != land[y][x-1] {
		results[self].Border++
	}
	if !inBounds(x+1, y, land) || land[y][x] != land[y][x+1] {
		results[self].Border++
	}

	visited[y][x] = true
}

func inBounds(x, y int, land [][]int) bool {
	return x >= 0 && y >= 0 && y < len(land) && x < len(land[y])
}

/// Part 2 \\\

func parseInput(lines []string) ([][]int, [][]bool) {
	land := make([][]rune, len(lines))
	visited := make([][]bool, len(lines))

	for y, line := range lines {
		land[y] = make([]rune, len(line))
		visited[y] = make([]bool, len(lines))
		for x, char := range line {
			land[y][x] = char
		}
	}

	// Flood the land to create uniquely identifiable regions.
	return floodLand(land), visited
}

func solvePart2(lines []string) (int, error) {
	results := make(map[int]*Result)
	regions, visited := parseInput(lines)

	// Count the area and borders of each region.
	for y, line := range regions {
		for x := range line {
			quantifyRegion(x, y, regions, visited, results)
		}
	}

	// Zoom the map so we can draw more precise boundary lines to count the sides.
	zoomed := zoom(regions)

	var sum int
	edges := make(map[int]int)
	for y, line := range zoomed {
		for x, id := range line {
			if id != 0 {
				edges[id] += countEdges(x, y, zoomed)
			}
		}
	}
	for id, result := range results {
		sum += result.Area * edges[id]
	}
	return sum, nil
}
