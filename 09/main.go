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

/// Part 1 \\\

func solve(lines []string) (int, error) {
	if len(lines) != 1 {
		return 0, fmt.Errorf("invalid input: too many lines")
	}

	var diskSizes []int
	var freeSpace []int
	for i, char := range lines[0] {
		size, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, err
		}
		if i%2 == 0 {
			diskSizes = append(diskSizes, size)
		} else {
			freeSpace = append(freeSpace, size)
		}
	}

	var str string
	var total, idx int
	sum := func(id, length int) {
		for ; length > 0; length-- {
			str += strconv.Itoa(id)
			total += id * idx
			idx++
		}
	}

	left, right := 0, len(diskSizes)-1
	for left <= right {
		// Left first
		sum(left, diskSizes[left])
		diskSizes[left] = 0

		if left == right {
			break
		}

		// Squeeze in right
		for freeSpace[left] > 0 {
			if diskSizes[right] == freeSpace[left] {
				sum(right, freeSpace[left])
				diskSizes[right] = 0
				right--
				break
			} else if diskSizes[right] > freeSpace[left] {
				sum(right, freeSpace[left])
				diskSizes[right] -= freeSpace[left]
				freeSpace[left] = 0
				break
			} else {
				sum(right, diskSizes[right])
				freeSpace[left] -= diskSizes[right]
				diskSizes[right] = 0
				right--
			}
		}
		left++
	}

	return total, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
