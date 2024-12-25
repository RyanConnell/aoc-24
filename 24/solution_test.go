package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay24(t *testing.T) {
	tester.TimeAndCheck(24, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    2024,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    42049478636360,
		},
	})
}
