package main

import (
	"fmt"
	"strings"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/sample2.txt")

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

const (
	NORTH = '^'
	EAST  = '>'
	SOUTH = 'v'
	WEST  = '<'
)

var offsets = map[rune]Point{
	NORTH: {Y: -1},
	EAST:  {X: 1},
	SOUTH: {Y: 1},
	WEST:  {X: -1},
}

type Point struct {
	X, Y int
}

func parseInput(lines []string) (*Point, [][]rune, string) {
	var directions []string
	var board [][]rune
	player := &Point{}
	for y, line := range lines {
		if line == "" {
			directions = lines[y+1:]
			break
		}
		board = append(board, make([]rune, len(line)))
		for x, char := range line {
			if char == '@' {
				player.X, player.Y = x, y
				board[y][x] = '@'
				continue
			}
			board[y][x] = char
		}
	}
	return player, board, strings.Join(directions, "")
}

func solve(lines []string) (int, error) {
	player, board, directions := parseInput(lines)

	for _, dir := range directions {
		if move(dir, player.X, player.Y, board, false) {
			move(dir, player.X, player.Y, board, true)
			offset := offsets[dir]
			player.X += offset.X
			player.Y += offset.Y
		}
	}

	var sum int
	for y, line := range board {
		for x, char := range line {
			if char == 'O' {
				sum += (y * 100) + x
			}
		}
	}

	return sum, nil
}

func move(dir rune, x, y int, board [][]rune, mutate bool) bool {
	if !inBounds(x, y, board) {
		return false
	}
	if board[y][x] == '.' {
		return true
	}
	nextX := x + offsets[dir].X
	nextY := y + offsets[dir].Y
	switch board[y][x] {
	case 'O', '@':
		if !move(dir, nextX, nextY, board, mutate) {
			return false
		}
		if mutate {
			board[y][x], board[y+offsets[dir].Y][x+offsets[dir].X] =
				board[y+offsets[dir].Y][x+offsets[dir].X], board[y][x]
		}
	case '[':
		// If moving left we only need to care if the left side moves
		switch dir {
		case NORTH, SOUTH:
			if !move(dir, nextX, nextY, board, mutate) || !move(dir, nextX+1, nextY, board, mutate) {
				return false
			}
			if mutate {
				board[y][x], board[y+offsets[dir].Y][x+offsets[dir].X] =
					board[y+offsets[dir].Y][x+offsets[dir].X], board[y][x]
				board[y][x+1], board[y+offsets[dir].Y][x+offsets[dir].X+1] =
					board[y+offsets[dir].Y][x+offsets[dir].X+1], board[y][x+1]
			}
		case EAST:
			if !move(dir, nextX+1, nextY, board, mutate) {
				return false
			}
			if mutate {
				board[y][x+2], board[y][x+1] = board[y][x+1], board[y][x+2]
				board[y][x+1], board[y][x] = board[y][x], board[y][x+1]
			}
		case WEST:
			if !move(dir, nextX, y, board, mutate) {
				return false
			}
			if mutate {
				board[y][x-1], board[y][x] = board[y][x], board[y][x-1]
				board[y][x+1], board[y][x] = board[y][x], board[y][x+1]
			}
		}
	case ']':
		switch dir {
		case NORTH, SOUTH:
			if !move(dir, nextX, nextY, board, mutate) || !move(dir, nextX-1, nextY, board, mutate) {
				return false
			}
			if mutate {
				board[y][x], board[y+offsets[dir].Y][x+offsets[dir].X] =
					board[y+offsets[dir].Y][x+offsets[dir].X], board[y][x]
				board[y][x-1], board[y+offsets[dir].Y][x+offsets[dir].X-1] =
					board[y+offsets[dir].Y][x+offsets[dir].X-1], board[y][x-1]
			}
		case WEST:
			if !move(dir, nextX-1, nextY, board, mutate) {
				return false
			}
			if mutate {
				board[y][x-2], board[y][x-1] = board[y][x-1], board[y][x-2]
				board[y][x-1], board[y][x] = board[y][x], board[y][x-1]
			}
		case EAST:
			if !move(dir, nextX, nextY, board, mutate) {
				return false
			}
			if mutate {
				board[y][x+1], board[y][x] = board[y][x], board[y][x+1]
				board[y][x-1], board[y][x] = board[y][x], board[y][x-1]
			}
		}
	}
	return true
}

func inBounds(x, y int, board [][]rune) bool {
	return x >= 0 && y >= 0 && y < len(board) && x < len(board[y]) && board[y][x] != '#'
}

func printBoard(player *Point, board [][]rune) {
	for y, line := range board {
		for x, char := range line {
			if player != nil && player.X == x && player.Y == y {
				fmt.Print("@")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

/// Part 2 \\\

func parseScaledInput(lines []string) (*Point, [][]rune, string) {
	var directions []string
	var board [][]rune
	player := &Point{}
	for y, line := range lines {
		if line == "" {
			directions = lines[y+1:]
			break
		}
		board = append(board, make([]rune, len(line)*2))
		for x, char := range line {
			switch char {
			case '#':
				board[y][(x*2)+0] = char
				board[y][(x*2)+1] = char
			case 'O':
				board[y][(x*2)+0] = '['
				board[y][(x*2)+1] = ']'
			case '.':
				board[y][(x*2)+0] = char
				board[y][(x*2)+1] = char
			case '@':
				player.X, player.Y = x*2, y
				board[y][(x*2)+0] = '@'
				board[y][(x*2)+1] = '.'
			}
		}
	}
	return player, board, strings.Join(directions, "")
}

func solvePart2(lines []string) (int, error) {
	player, board, directions := parseScaledInput(lines)

	for _, dir := range directions {
		if move(dir, player.X, player.Y, board, false) {
			move(dir, player.X, player.Y, board, true)
			offset := offsets[dir]
			player.X += offset.X
			player.Y += offset.Y
		}
	}

	var sum int
	for y, line := range board {
		for x, char := range line {
			if char == '[' {
				sum += (y * 100) + x
			}
		}
	}

	return sum, nil
}
