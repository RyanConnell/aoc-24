package main

import (
	"fmt"

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

type Block struct {
	ID     int
	Length int
}

func solve(lines []string) (int, error) {
	if len(lines) != 1 {
		return 0, fmt.Errorf("invalid input: too many lines")
	}
	disk := parseInput(lines[0])

	left, right := 0, len(disk)-1
	for left <= right {
		if disk[left].ID != -1 {
			left++
			continue
		}
		if disk[right].ID == -1 {
			right--
			continue
		}

		if disk[left].Length == disk[right].Length {
			// Same size blocks, so we can just swap.
			disk[left], disk[right] = disk[right], disk[left]
		} else if disk[left].Length < disk[right].Length {
			// Not enough space to fit the entire block.
			disk[left].ID = disk[right].ID
			disk[right].Length -= disk[left].Length
		} else {
			// Extra space left after moving the entire block.
			remainder := &Block{ID: -1, Length: disk[left].Length - disk[right].Length}
			disk[left] = disk[right]
			disk[right] = &Block{ID: -1, Length: 0}
			disk = append(disk[:left+1], append([]*Block{remainder}, disk[left+1:]...)...)
		}
	}

	var sum, position int
	for _, block := range disk {
		if block.ID == -1 {
			continue
		}
		for i := 0; i < block.Length; i++ {
			sum += block.ID * position
			position++
		}
	}

	return sum, nil
}

func parseInput(line string) []*Block {
	var disk []*Block
	for i, char := range line {
		if i%2 == 0 {
			disk = append(disk, &Block{ID: i / 2, Length: int(char - '0')})
		} else {
			disk = append(disk, &Block{ID: -1, Length: int(char - '0')})
		}
	}
	return disk
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	if len(lines) != 1 {
		return 0, fmt.Errorf("invalid input: too many lines")
	}
	disk := parseInput(lines[0])

	for right := len(disk) - 1; right >= 0; right-- {
		if disk[right].ID == -1 {
			continue
		}
		for left := 0; left < right; left++ {
			if disk[left].ID != -1 {
				continue
			}
			if disk[left].Length == disk[right].Length {
				disk[right], disk[left] = disk[left], disk[right]
				break
			}
			if disk[left].Length > disk[right].Length {
				remainder := &Block{ID: -1, Length: disk[left].Length - disk[right].Length}
				disk[left].Length = disk[right].Length
				disk[left].ID, disk[right].ID = disk[right].ID, -1
				disk = append(disk[:left+1], append([]*Block{remainder}, disk[left+1:]...)...)
				break
			}
		}
	}

	var sum, position int
	for _, block := range disk {
		if block.ID == -1 {
			position += block.Length
			continue
		}
		for i := 0; i < block.Length; i++ {
			sum += block.ID * position
			position++
		}
	}

	return sum, nil
}
