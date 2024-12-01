package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay01(t *testing.T) {
	tester.TimeAndCheck(1, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    11,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    31,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    2285373,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    21142653,
		},
	})
}
