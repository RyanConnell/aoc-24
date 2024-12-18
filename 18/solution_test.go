package main

import (
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay18(t *testing.T) {
	tester.TimeAndCheck(18, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution: func(lines []string) (int, error) {
				return solve(lines, sampleScale, sampleBytes)
			},
			Expected: 22,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution: func(lines []string) (int, error) {
				return solve(lines, inputScale, inputBytes)
			},
			Expected: 290,
		},
	})
	tester.TimeAndCheck(18, []tester.TestCase[string]{
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution: func(lines []string) (string, error) {
				return solvePart2(lines, sampleScale, sampleBytes)
			},
			Expected: "6,1",
		},
		{
			Description: "Part 2 (input)",
			File:        "input.txt",
			Solution: func(lines []string) (string, error) {
				return solvePart2(lines, inputScale, inputBytes)
			},
			Expected: "64,54",
		},
	})
}
