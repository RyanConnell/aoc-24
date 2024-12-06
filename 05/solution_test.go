package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay05(t *testing.T) {
	tester.TimeAndCheck(5, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    143,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    5964,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    123,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    4719,
		},
	})
}
