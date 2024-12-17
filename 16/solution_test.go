package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay15(t *testing.T) {
	tester.TimeAndCheck(15, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    7036,
		},
		{
			Description: "Part 1 (sample2)",
			File:        "sample2.txt",
			Solution:    solve,
			Expected:    11048,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    105496,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    45,
		},
		{
			Description: "Part 2 (sample2)",
			File:        "sample2.txt",
			Solution:    solvePart2,
			Expected:    64,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    524,
		},
	})
}
