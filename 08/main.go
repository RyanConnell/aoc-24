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
	//fmt.Println(p, target, xDistance, yDistance)
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

func (p Point) String() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

/// Part 1 \\\

func solve(lines []string) (int, error) {
	nodesPerChar := make(map[rune][]Point)
	antinodes := make(map[string]Point)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				nodesPerChar[char] = append(nodesPerChar[char], Point{X: x, Y: y})
			}
		}
	}

	for _, nodes := range nodesPerChar {
		for i, node := range nodes {
			for _, target := range nodes[i+1:] {
				for _, anode := range node.Antinodes(target) {
					if anode.X < 0 || anode.Y < 0 || anode.X >= len(lines[0]) || anode.Y >= len(lines) {
						continue
					}
					antinodes[anode.String()] = anode
				}
			}
		}
	}

	return len(antinodes), nil
}

func printMap(lines []string, antinodes map[string]Point) {
	m := make([][]rune, len(lines))
	for y, line := range lines {
		m[y] = make([]rune, len(line))
		for x, char := range line {
			m[y][x] = char
		}
	}

	for _, point := range antinodes {
		m[point.Y][point.X] = '#'
	}
	for _, line := range m {
		for _, char := range line {
			fmt.Print(string(char), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
