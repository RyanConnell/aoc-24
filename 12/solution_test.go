package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay12(t *testing.T) {
	tester.TimeAndCheck(12, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    140,
		},
		{
			Description: "Part 1 (sample2)",
			File:        "sample2.txt",
			Solution:    solve,
			Expected:    1930,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    1431440,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    80,
		},
		{
			Description: "Part 2 (sample2)",
			File:        "sample2.txt",
			Solution:    solvePart2,
			Expected:    1206,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    869070,
		},
	})
}
