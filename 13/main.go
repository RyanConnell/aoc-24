package main

import (
	"fmt"
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

type Point struct {
	X, Y int
}

type Game struct {
	buttons []Point
	prize   Point
}

// evaluateButtons returns the number of times we need to press button A and B in order to get to
// our target. Note that this is a best estimate (we're rounding to an int afterwards) so you
// should run the 'verify' function afterwards to ensure it is accurate.
func (g Game) evaluateButtons() (int, int) {
	x1, x2, x3 := g.buttons[0].X, g.buttons[1].X, g.prize.X
	y1, y2, y3 := g.buttons[0].Y, g.buttons[1].Y, g.prize.Y

	right := (x3 * y2) - (x2 * y3)
	denominator := x1*y2 - x2*y1

	aCount := right / denominator
	bCount := (y3 - (y1 * aCount)) / y2

	return aCount, bCount
}

func (g Game) verify(aCount, bCount int) bool {
	x := (aCount * g.buttons[0].X) + (bCount * g.buttons[1].X)
	y := (aCount * g.buttons[0].Y) + (bCount * g.buttons[1].Y)
	return x == g.prize.X && y == g.prize.Y
}

func parseInput(lines []string) ([]Game, error) {
	var games []Game
	var gameID int
	for _, line := range lines {
		if len(games) == gameID {
			games = append(games, Game{})
		}
		if strings.HasPrefix(line, "Button") {
			vals := strings.Split(line[9:], ", ")
			point, err := parseCoords(vals[0][2:], vals[1][2:])
			if err != nil {
				return nil, err
			}
			games[gameID].buttons = append(games[gameID].buttons, point)
		} else if strings.HasPrefix(line, "Prize") {
			vals := strings.Split(line[7:], ", ")
			point, err := parseCoords(vals[0][2:], vals[1][2:])
			if err != nil {
				return nil, err
			}
			games[gameID].prize = point
			gameID++
		}
	}
	return games, nil
}

func parseCoords(xs, ys string) (Point, error) {
	x, err := strconv.Atoi(xs)
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.Atoi(ys)
	if err != nil {
		return Point{}, err
	}
	return Point{x, y}, nil
}

func solve(lines []string) (int, error) {
	games, err := parseInput(lines)
	if err != nil {
		return 0, err
	}
	var result int
	for _, game := range games {
		a, b := game.evaluateButtons()
		if game.verify(a, b) {
			cost := (a * 3) + b
			result += cost
		}
	}
	return result, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	games, err := parseInput(lines)
	if err != nil {
		return 0, err
	}
	var result int
	for _, game := range games {
		game.prize.X += 10000000000000
		game.prize.Y += 10000000000000
		a, b := game.evaluateButtons()
		if game.verify(a, b) {
			cost := (a * 3) + b
			result += cost
		}
	}
	return result, nil
}
