package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay06(t *testing.T) {
	tester.TimeAndCheck(6, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    41,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    5153,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    6,
		},
		{
			Description: "Part 2 (custom)",
			File:        "custom.txt",
			Solution:    solvePart2,
			Expected:    3,
		},
		{
			Description: "Part 2 (custom2)",
			File:        "custom2.txt",
			Solution:    solvePart2,
			Expected:    1,
		},
		{
			Description: "Part 2 (custom3)",
			File:        "custom3.txt",
			Solution:    solvePart2,
			Expected:    3,
		},
		{
			Description: "Part 2 (custom4)",
			File:        "custom4.txt",
			Solution:    solvePart2,
			Expected:    1,
		},
		{
			Description: "Part 2 (custom5)",
			File:        "custom5.txt",
			Solution:    solvePart2,
			Expected:    1,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    1711,
		},
	})
}
