package main

import (
	"fmt"
	"strconv"

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

type Point struct {
	X, Y int
}

func (p Point) Antinodes(target Point) []Point {
	xDistance := float64(p.X) - float64(target.X)
	yDistance := float64(p.Y) - float64(target.Y)
	return []Point{
		{
			X: int(p.X + int(xDistance)),
			Y: int(p.Y + int(yDistance)),
		},
		{
			X: int(target.X - int(xDistance)),
			Y: int(target.Y - int(yDistance)),
		},
	}
}

func (p Point) HarmonicAntinodes(target Point, bounds *Point) []Point {
	xDistance := float64(p.X) - float64(target.X)
	yDistance := float64(p.Y) - float64(target.Y)
	var antinodes []Point
	for i := 0; ; i++ {
		var added bool
		pAdj := Point{
			X: int(p.X + int(xDistance)*i),
			Y: int(p.Y + int(yDistance)*i),
		}
		tAdj := Point{
			X: int(target.X - int(xDistance)*i),
			Y: int(target.Y - int(yDistance)*i),
		}
		if bounds == nil {
			return []Point{pAdj, tAdj}
		}
		if pAdj.X < bounds.X && pAdj.Y < bounds.Y && pAdj.X >= 0 && pAdj.Y >= 0 {
			antinodes = append(antinodes, pAdj)
			added = true
		}
		if tAdj.X < bounds.X && tAdj.Y < bounds.Y && tAdj.X >= 0 && tAdj.Y >= 0 {
			antinodes = append(antinodes, tAdj)
			added = true
		}
		if !added {
			break
		}
	}
	return antinodes
}

func (p Point) String() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

/// Part 1 \\\

func solve(lines []string) (int, error) {
	return solution(lines, false)
}

func solution(lines []string, allowHarmonics bool) (int, error) {
	nodesPerChar := make(map[rune][]Point)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				nodesPerChar[char] = append(nodesPerChar[char], Point{X: x, Y: y})
			}
		}
	}

	antinodeSet := make(map[string]Point)
	for _, nodes := range nodesPerChar {
		for i, node := range nodes {
			for _, target := range nodes[i+1:] {
				var antinodes []Point
				if allowHarmonics {
					antinodes = node.HarmonicAntinodes(target,
						&Point{X: len(lines[0]), Y: len(lines)})
				} else {
					antinodes = node.Antinodes(target)
				}
				for _, anode := range antinodes {
					if anode.X < 0 || anode.Y < 0 || anode.X >= len(lines[0]) || anode.Y >= len(lines) {
						continue
					}
					antinodeSet[anode.String()] = anode
				}
			}
		}
	}

	return len(antinodeSet), nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return solution(lines, true)
}
