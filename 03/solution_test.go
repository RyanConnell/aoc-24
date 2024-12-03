package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay03(t *testing.T) {
	tester.TimeAndCheck(3, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    161,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    155955228,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample2.txt",
			Solution:    solvePart2,
			Expected:    48,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    100189366,
		},
	})
}
