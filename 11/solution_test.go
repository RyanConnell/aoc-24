package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay00(t *testing.T) {
	tester.TimeAndCheck(0, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    55312,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    229043,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    272673043446478,
		},
	})
}
