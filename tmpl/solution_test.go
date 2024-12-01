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
			Expected:    0,
		},
	})
}
