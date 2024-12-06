package main

import (
	"fmt"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/custom4.txt")

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

type Point struct {
	X, Y                   int
	Solid, Start, Loopable bool
	VisitedDirections      []int
}

func solve(lines []string) (int, error) {
	var guard *Point
	board := make([][]rune, len(lines))
	for y, line := range lines {
		board[y] = make([]rune, len(line))
		for x, char := range line {
			if char == '#' {
				board[y][x] = char
			}
			if char == '^' {
				guard = &Point{X: x, Y: y}
				board[y][x] = 'X'
			}
		}
	}

	direction := NORTH
	for move(guard, direction, board) {
		direction = (direction + 1) % 4
	}

	var count int
	for _, line := range board {
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}

	return count, nil
}

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

// move returns true if there are more moves to make
func move(guard *Point, direction int, board [][]rune) bool {
	next := Point{X: guard.X, Y: guard.Y}
	for {
		switch direction {
		case NORTH:
			next.Y--
		case EAST:
			next.X++
		case SOUTH:
			next.Y++
		case WEST:
			next.X--
		}

		if next.Y < 0 || next.X < 0 || next.Y >= len(board) || next.X >= len(board[next.Y]) {
			return false
		}
		if board[next.Y][next.X] == '#' {
			break
		}
		board[next.Y][next.X] = 'X'
		guard.X, guard.Y = next.X, next.Y
	}
	return true
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	var guard *Point
	pointMap := make([][]*Point, len(lines))
	for y, line := range lines {
		pointMap[y] = make([]*Point, len(line))
		for x, char := range line {
			pointMap[y][x] = &Point{X: x, Y: y}
			if char == '#' {
				pointMap[y][x].Solid = true
			} else if char == '^' {
				guard = &Point{X: x, Y: y}
				pointMap[y][x].Start = true
				pointMap[y][x].VisitedDirections = []int{NORTH}
			}
		}
	}

	board := NewBoard(pointMap)

	direction := NORTH
	for movePart2(guard, direction, board, false) {
		direction = (direction + 1) % 4
		board.Points[guard.Y][guard.X].VisitedDirections = append(
			board.Points[guard.Y][guard.X].VisitedDirections, direction)
	}

	var loopCount int
	for _, line := range board.Points {
		for _, point := range line {
			if point.Loopable && !point.Solid && !point.Start {
				loopCount++
			}
		}
	}

	return loopCount, nil
}

type Board struct {
	Points [][]*Point
}

func NewBoard(points [][]*Point) *Board {
	return &Board{Points: points}
}

func (p *Point) Next(direction int) Point {
	next := Point{X: p.X, Y: p.Y}
	switch direction {
	case NORTH:
		next.Y--
	case EAST:
		next.X++
	case SOUTH:
		next.Y++
	case WEST:
		next.X--
	}
	return next
}

func (b *Board) CheckLoopable(p Point, direction int) bool {
	key := func(p Point, direction int) string {
		return fmt.Sprintf("%d-%d-%d", p.X, p.Y, direction)
	}
	visited := map[string]struct{}{}
	last, next := p, p.Next(direction)
	for {
		if _, ok := visited[key(next, direction)]; ok {
			return true
		}
		if next.Y < 0 || next.X < 0 {
			return false
		}
		if next.Y >= len(b.Points) || next.X >= len(b.Points[next.Y]) {
			return false
		}
		if b.Points[next.Y][next.X].Solid {
			direction = (direction + 1) % 4
			next = last
			continue
		}
		for _, dir := range b.Points[next.Y][next.X].VisitedDirections {
			if dir == direction {
				return true
			}
		}
		visited[key(next, direction)] = struct{}{}
		last, next = next, next.Next(direction)
	}
}

// move returns true if there are more moves to make
func movePart2(guard *Point, direction int, board *Board, check bool) bool {
	for {
		next := guard.Next(direction)
		if next.Y < 0 || next.X < 0 {
			return false
		}
		if next.Y >= len(board.Points) || next.X >= len(board.Points[next.Y]) {
			return false
		}
		if board.Points[next.Y][next.X].Solid {
			break
		}

		// Ensure that we haven't travelled through this point before.
		if len(board.Points[next.Y][next.X].VisitedDirections) == 0 {
			if !board.Points[next.Y][next.X].Start {
				if board.CheckLoopable(*guard, (direction+1)%4) {
					board.Points[next.Y][next.X].Loopable = true
				}
			}
		}

		board.Points[next.Y][next.X].VisitedDirections = append(
			board.Points[next.Y][next.X].VisitedDirections, direction)
		guard.X, guard.Y = next.X, next.Y

	}
	return true
}

func printMap(board [][]*Point) {
	for y := range board {
		for x := range board[y] {
			if board[y][x].Solid {
				fmt.Print("O ")
			} else if board[y][x].Start {
				fmt.Print("S ")
			} else if board[y][x].Loopable {
				fmt.Print("+ ")
			} else if len(board[y][x].VisitedDirections) != 0 {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
