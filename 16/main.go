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

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

var offsets = map[int]Point{
	NORTH: {Y: -1},
	EAST:  {X: 1},
	SOUTH: {Y: 1},
	WEST:  {X: -1},
}

type Point struct {
	X, Y             int
	Tile             rune
	Best             bool
	scoreToDirection map[int]int
}

func (p *Point) Step(direction int) Point {
	return Point{X: p.X + offsets[direction].X, Y: p.Y + offsets[direction].Y}
}

func (p *Point) InBounds(board [][]*Point) bool {
	return p.X > 0 && p.Y > 0 && p.Y < len(board) && p.Y < len(board[p.Y])
}

func navigate(direction, score int, pos Point, board [][]*Point) {
	p := board[pos.Y][pos.X]
	if bestScore, ok := p.scoreToDirection[direction]; ok {
		if score >= bestScore {
			return
		}
		p.scoreToDirection[direction] = score
	}
	if p.Tile == 'E' {
		// Duplicate scores at end for simplified viewing.
		p.scoreToDirection[direction] = score
		return
	}

	// Check direction
	next := pos.Step(direction)
	if next.InBounds(board) && board[next.Y][next.X].Tile != '#' {
		if val, ok := p.scoreToDirection[direction]; !ok || val < score+1 {
			p.scoreToDirection[direction] = score + 1
			navigate(direction, score+1, next, board)
		}
	}

	// Check direction -1
	next = pos.Step((direction + 3) % 4)
	if next.InBounds(board) && board[next.Y][next.X].Tile != '#' {
		if val, ok := p.scoreToDirection[(direction+3)%4]; !ok || val > score+1001 {
			p.scoreToDirection[(direction+3)%4] = score + 1001
			navigate((direction+3)%4, score+1001, next, board)
		}
	}

	// Check direction +1
	next = pos.Step((direction + 1) % 4)
	if next.InBounds(board) && board[next.Y][next.X].Tile != '#' {
		if val, ok := p.scoreToDirection[(direction+1)%4]; !ok || val > score+1001 {
			p.scoreToDirection[(direction+1)%4] = score + 1001
			navigate((direction+1)%4, score+1001, next, board)
		}
	}
}

func countTiles(score int, pos Point, board [][]*Point) int {
	if board[pos.Y][pos.X].Tile == 'S' {
		board[pos.Y][pos.X].Best = true
		return 1
	}
	if board[pos.Y][pos.X].Best {
		return 0
	}
	var count int
	for i := 0; i < 4; i++ {
		next := pos.Step(i)
		nextScore := board[next.Y][next.X].scoreToDirection[(i+2)%4]
		if nextScore != -1 && (nextScore == score-1 || nextScore == score-1001) {
			count += countTiles(nextScore, next, board)
		}
	}
	board[pos.Y][pos.X].Best = true
	return count + 1
}

func solve(lines []string) (int, error) {
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
			board[y][x] = &Point{
				X:                x,
				Y:                y,
				Tile:             char,
				scoreToDirection: make(map[int]int),
			}
		}
	}

	navigate(EAST, 0, start, board)
	bestScore := -1
	for _, score := range board[end.Y][end.X].scoreToDirection {
		if bestScore == -1 || score < bestScore {
			bestScore = score
		}
	}
	return bestScore, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
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
			board[y][x] = &Point{
				X:                x,
				Y:                y,
				Tile:             char,
				scoreToDirection: make(map[int]int),
			}
		}
	}

	navigate(EAST, 0, start, board)
	bestScore := -1
	bestDir := -1
	for dir, score := range board[end.Y][end.X].scoreToDirection {
		if bestScore == -1 || score < bestScore {
			bestScore = score
			bestDir = dir
		}
	}
	count := countTiles(bestScore, end.Step((bestDir+2)%4), board)
	return count + 1, nil
}
