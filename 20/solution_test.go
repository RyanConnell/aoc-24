package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay20(t *testing.T) {
	tester.TimeAndCheck(20, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 20, 2) },
			Expected:    5,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 100, 2) },
			Expected:    1499,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 50, 20) },
			Expected:    285,
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 100, 20) },
			Expected:    1027164,
		},
	})
}
