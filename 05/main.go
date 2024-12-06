package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-24/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

func parseInput(lines []string) ([]string, []string) {
	for i, line := range lines {
		if line == "" {
			return lines[:i], lines[i+1:]
		}
	}
	return nil, nil
}

func solve(lines []string) (int, error) {
	pageRules, pageOrders := parseInput(lines)

	pageRuleMap := make(map[string][]string)
	for _, pageRule := range pageRules {
		parts := strings.Split(pageRule, "|")
		if _, ok := pageRuleMap[parts[1]]; ok {
			pageRuleMap[parts[1]] = append(pageRuleMap[parts[1]], parts[0])
		} else {
			pageRuleMap[parts[1]] = []string{parts[0]}
		}
	}

	var sum int
	for _, line := range pageOrders {
		if !checkOrder(line, pageRuleMap) {
			continue
		}
		parts := strings.Split(line, ",")
		middleStr := parts[len(parts)/2]
		middle, err := strconv.Atoi(middleStr)
		if err != nil {
			return 0, err
		}
		sum += middle
	}

	return sum, nil
}

func checkOrder(orderStr string, rules map[string][]string) bool {
	seenPage := make(map[string]struct{})
	hasPage := make(map[string]struct{})
	order := strings.Split(orderStr, ",")
	for _, p := range order {
		hasPage[p] = struct{}{}
	}
	for _, p := range order {
		for _, required := range rules[p] {
			if _, ok := hasPage[required]; !ok {
				continue
			}
			if _, ok := seenPage[required]; !ok {
				return false
			}
		}
		seenPage[p] = struct{}{}
	}
	return true
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	pageRules, pageOrders := parseInput(lines)

	pageRuleMap := make(map[string][]string)
	for _, pageRule := range pageRules {
		parts := strings.Split(pageRule, "|")
		if _, ok := pageRuleMap[parts[1]]; ok {
			pageRuleMap[parts[1]] = append(pageRuleMap[parts[1]], parts[0])
		} else {
			pageRuleMap[parts[1]] = []string{parts[0]}
		}
	}

	var sum int
	for _, line := range pageOrders {
		if checkOrder(line, pageRuleMap) {
			continue
		}
		line = fixOrder(line, pageRuleMap)

		parts := strings.Split(line, ",")
		middleStr := parts[len(parts)/2]
		middle, err := strconv.Atoi(middleStr)
		if err != nil {
			return 0, err
		}
		sum += middle
	}

	return sum, nil
}

func fixOrder(orderStr string, rules map[string][]string) string {
	order := strings.Split(orderStr, ",")
	seenPage := make(map[string]struct{})
	hasPage := make(map[string]struct{})
	for _, p := range order {
		hasPage[p] = struct{}{}
	}
	var i int
	for i < len(order) {
		var added bool
		for _, required := range rules[order[i]] {
			if _, ok := hasPage[required]; !ok {
				continue
			}
			if _, ok := seenPage[required]; !ok {
				var order2 []string
				order2 = append(order2, order[:i]...)
				order2 = append(order2, required)
				order2 = append(order2, order[i:]...)
				order = order2
				added = true
				break
			}
		}
		if added {
			continue
		}

		seenPage[order[i]] = struct{}{}
		i++
	}

	seen := make(map[string]struct{})
	for i := 0; i < len(order); i++ {
		if _, ok := seen[order[i]]; !ok {
			seen[order[i]] = struct{}{}
			continue
		}
		order = append(order[:i], order[i+1:]...)
		i--
	}
	return strings.Join(order, ",")
}

func checkOrderPart2(orderStr string, rules map[string][]string) bool {
	seenPage := make(map[string]struct{})
	hasPage := make(map[string]struct{})
	order := strings.Split(orderStr, ",")
	for _, p := range order {
		hasPage[p] = struct{}{}
	}
	for _, p := range order {
		for _, required := range rules[p] {
			if _, ok := hasPage[required]; !ok {
				continue
			}
			if _, ok := seenPage[required]; !ok {
				return false
			}
		}
		seenPage[p] = struct{}{}
	}
	return true
}
