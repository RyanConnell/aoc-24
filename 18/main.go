package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

const (
	sampleScale = 7
	inputScale  = 71

	sampleBytes = 12
	inputBytes  = 1024
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines, inputScale, inputBytes)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solvePart2(lines, inputScale, inputBytes)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %s\n", solutionPart2)
}

type Point struct {
	X     int
	Y     int
	Solid bool
	Score int
}

/// Part 1 \\\

func solve(lines []string, scale, bytes int) (int, error) {
	board := make([][]*Point, scale)
	for y := range board {
		board[y] = make([]*Point, scale)
		for x := range board[y] {
			board[y][x] = &Point{X: x, Y: y, Score: -1}
		}
	}

	for i := 0; i < bytes; i++ {
		vals := strings.Split(lines[i], ",")
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(vals[1])
		if err != nil {
			return 0, err
		}
		board[y][x].Solid = true
	}

	flood(0, 0, 0, board)
	return board[scale-1][scale-1].Score, nil
}

func inBounds(x, y int, board [][]*Point) bool {
	return x >= 0 && y >= 0 && y < len(board) && x < len(board[y]) && !board[y][x].Solid
}

func flood(score int, x, y int, board [][]*Point) {
	if !inBounds(x, y, board) {
		return
	}
	if board[y][x].Score != -1 && board[y][x].Score <= score {
		return
	}
	board[y][x].Score = score

	flood(score+1, x+1, y, board)
	flood(score+1, x-1, y, board)
	flood(score+1, x, y+1, board)
	flood(score+1, x, y-1, board)
}

/// Part 2 \\\

func prepareBoard(board [][]*Point, instructions []Point, midpoint int) {
	for _, line := range board {
		for _, p := range line {
			p.Solid = false
			p.Score = -1
		}
	}
	for i := 0; i <= midpoint; i++ {
		board[instructions[i].Y][instructions[i].X].Solid = true
	}
}

func solvePart2(lines []string, scale, bytes int) (string, error) {
	board := make([][]*Point, scale)
	for y := range board {
		board[y] = make([]*Point, scale)
		for x := range board[y] {
			board[y][x] = &Point{X: x, Y: y, Score: -1}
		}
	}

	instructions := make([]Point, len(lines))
	for i := 0; i < len(lines); i++ {
		vals := strings.Split(lines[i], ",")
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(vals[1])
		if err != nil {
			return "", err
		}
		instructions[i].X = x
		instructions[i].Y = y
	}

	left, right := 0, len(lines)-1
	for {
		midpoint := left + (right-left)/2
		prepareBoard(board, instructions, midpoint)
		flood(0, 0, 0, board)
		steps := board[scale-1][scale-1].Score
		if steps != -1 {
			// If we're still valid, we go further right.
			left = midpoint
		} else {
			// If we're no longer valid we go further left.
			right = midpoint
		}

		if right-left == 1 {
			return fmt.Sprintf("%d,%d", instructions[midpoint].X, instructions[midpoint].Y), nil
		}
	}
}
