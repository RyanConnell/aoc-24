package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay19(t *testing.T) {
	tester.TimeAndCheck(19, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    6,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    283,
		},
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    16,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    615388132411142,
		},
	})
}
