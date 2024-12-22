package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay22(t *testing.T) {
	tester.TimeAndCheck(22, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    37327623,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    15335183969,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    24,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    1696,
		},
	})
}
