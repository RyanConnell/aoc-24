package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay13(t *testing.T) {
	tester.TimeAndCheck(13, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    480,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    28753,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    875318608908,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    102718967795500,
		},
	})
}
