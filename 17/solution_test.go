package main

import (
	"reflect"
	"testing"

	"github.com/RyanConnell/aoc-24/pkg/tester"
)

func TestDay17(t *testing.T) {
	tester.TimeAndCheck(17, []tester.TestCase[string]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    "4,6,3,5,6,3,5,2,1,0",
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    "6,7,5,2,1,3,5,1,7",
		},
	})
}

func TestMemory(t *testing.T) {
	testCases := map[string]struct {
		registers         []int
		instructions      []int
		expectedRegisters []int
		expected          string
	}{
		"bst example": {
			registers:         []int{0, 0, 9},
			instructions:      []int{2, 6},
			expectedRegisters: []int{0, 1, 9},
			expected:          "",
		},
		"out example": {
			registers:         []int{10, 0, 0},
			instructions:      []int{5, 0, 5, 1, 5, 4},
			expectedRegisters: []int{10, 0, 0},
			expected:          "0,1,2",
		},
		"adv example": {
			registers:         []int{2024, 0, 0},
			instructions:      []int{0, 1, 5, 4, 3, 0},
			expectedRegisters: []int{0, 0, 0},
			expected:          "4,2,5,6,7,7,7,7,3,1,0",
		},
		"bdv example": {
			registers:         []int{2024, 0, 0},
			instructions:      []int{6, 1, 5, 5},
			expectedRegisters: []int{2024, 1012, 0},
			expected:          "4",
		},
		"cdv example": {
			registers:         []int{2024, 0, 0},
			instructions:      []int{7, 1, 5, 6},
			expectedRegisters: []int{2024, 0, 1012},
			expected:          "4",
		},
		"bxl example": {
			registers:         []int{0, 29, 0},
			instructions:      []int{1, 7},
			expectedRegisters: []int{0, 26, 0},
			expected:          "",
		},
		"bxc example": {
			registers:         []int{0, 2024, 43690},
			instructions:      []int{4, 0},
			expectedRegisters: []int{0, 44354, 43690},
			expected:          "",
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			memory := &Memory{
				Registers:    tc.registers,
				Instructions: tc.instructions,
			}
			out := memory.run()
			if tc.expected != out {
				t.Errorf("output did not match; want %q got %q", tc.expected, out)
			}
			if !reflect.DeepEqual(tc.expectedRegisters, memory.Registers) {
				t.Errorf("registers did not match; want %q got %q",
					tc.expectedRegisters, memory.Registers)
			}
		})
	}
}
