package main

import (
	"fmt"
	"math"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines, 100, 2)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, 100, 20)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

type Point struct {
	X    int
	Y    int
	Tile rune
	Cost int
}

func parseInput(lines []string) (Point, Point, [][]*Point) {
	board := make([][]*Point, len(lines))
	var start, end Point
	for y, line := range lines {
		board[y] = make([]*Point, len(line))
		for x, char := range line {
			if char == 'S' {
				start = Point{X: x, Y: y}
			}
			if char == 'E' {
				end = Point{X: x, Y: y}
			}
			board[y][x] = &Point{X: x, Y: y, Tile: char, Cost: math.MaxInt64}
		}
	}
	return start, end, board
}

func solve(lines []string, theshold int, radius int) (int, error) {
	start, end, board := parseInput(lines)
	flood(0, end.X, end.Y, board)
	cheats := traverse(theshold, start.X, start.Y, board, generatePositions(radius))
	return cheats, nil
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func generatePositions(radius int) []Point {
	pointMap := make([][]string, (radius*2)+1)
	var width int
	var points []Point
	for y := -radius; y <= radius; y++ {
		pointMap[y+radius] = make([]string, (radius*2)+1)
		for x := -radius; x <= radius; x++ {
			pointMap[y+radius][x+radius] = "."
			if x == 0 && y == 0 {
				continue
			}
			if x <= width && x >= -width {
				cost := absInt(x) + absInt(y)
				points = append(points, Point{X: x, Y: y, Cost: cost})
				pointMap[y+radius][x+radius] = "X"
			}
		}
		if y < 0 {
			width++
		} else if y >= 0 {
			width--
		}
	}

	return points
}

var positionsToMove = []Point{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}

func traverse(threshold, x, y int, board [][]*Point, positionsToCheck []Point) int {
	// Check how many possible cheats we have from our current position.
	var cheats int
	for _, pos := range positionsToCheck {
		nextX, nextY := x+pos.X, y+pos.Y
		// todo: Only cheat if score <= 100
		if inBounds(nextX, nextY, board) && board[nextY][nextX].Cost < board[y][x].Cost-pos.Cost {
			savings := board[y][x].Cost - board[nextY][nextX].Cost - pos.Cost
			if savings >= threshold {
				//fmt.Println("Cheat: ", savings)
				cheats++
			}
		}
	}

	// Move to next spot in map.
	for _, pos := range positionsToMove {
		nextX, nextY := x+pos.X, y+pos.Y
		// todo: Only cheat if score <= 100
		if inBounds(nextX, nextY, board) && board[nextY][nextX].Cost == board[y][x].Cost-1 {
			cheats += traverse(threshold, nextX, nextY, board, positionsToCheck)
		}
	}

	return cheats
}

func inBounds(x, y int, board [][]*Point) bool {
	return y >= 0 && x >= 0 && y < len(board) && x < len(board[y]) && board[y][x].Tile != '#'
}

func flood(cost, x, y int, board [][]*Point) {
	if !inBounds(x, y, board) {
		return
	}
	if board[y][x].Cost < cost {
		return
	}
	board[y][x].Cost = cost

	flood(cost+1, x+1, y, board)
	flood(cost+1, x-1, y, board)
	flood(cost+1, x, y+1, board)
	flood(cost+1, x, y-1, board)
}
