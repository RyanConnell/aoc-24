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

/// Part 1 \\\

func solve(lines []string) (int, error) {
	var count int
	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {
				count += searchFromPos(x, y, lines)
			}
		}
	}
	return count, nil
}

var xyOffsets = [][]int{
	{1, 0},  // Right
	{-1, 0}, // Left
	{0, -1}, // Up
	{0, 1},  // Down
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func searchFromPos(x, y int, board []string) int {
	var count int
	for _, offset := range xyOffsets {
		if len(board) <= y+offset[1]*3 || y+offset[1]*3 < 0 {
			continue
		}
		if len(board[y]) <= x+offset[0]*3 || x+offset[0]*3 < 0 {
			continue
		}
		word := string([]byte{
			board[y][x],
			board[y+(offset[1]*1)][x+(offset[0]*1)],
			board[y+(offset[1]*2)][x+(offset[0]*2)],
			board[y+(offset[1]*3)][x+(offset[0]*3)],
		})
		if word == "XMAS" {
			count++
		}
	}
	return count
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	var count int
	for y, line := range lines {
		for x, char := range line {
			if char == 'A' && searchFromPosP2(x, y, lines) {
				count++
			}
		}
	}
	return count, nil
}

var validXMasPositions = [][]string{
	{"MM", "SS"},
	{"SS", "MM"},
	{"SM", "SM"},
	{"MS", "MS"},
}

func searchFromPosP2(x, y int, board []string) bool {
	if len(board) <= y+1 || y-1 < 0 {
		return false
	}
	if len(board[y]) <= x+1 || x-1 < 0 {
		return false
	}
	words := []string{
		string([]byte{board[y-1][x-1], board[y-1][x+1]}),
		string([]byte{board[y+1][x-1], board[y+1][x+1]}),
	}

	for _, validPos := range validXMasPositions {
		if words[0] == validPos[0] && words[1] == validPos[1] {
			return true
		}
	}
	return false
}
