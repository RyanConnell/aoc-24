package main

import (
	"fmt"
	"strconv"

	"github.com/RyanConnell/aoc-24/pkg/parser"
	"github.com/RyanConnell/aoc-24/pkg/pprof"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	pprof.MaybeProfile()
	defer pprof.MaybeStopProfile()

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

func solve(lines []string) (int, error) {
	var sum int
	for _, line := range lines {
		secret, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		result, _ := iterate(secret, 2000)
		sum += result
	}
	return sum, nil
}

func mix(a, b int) int { return a ^ b }

func prune(a int) int { return a % 16777216 }

func cost(n int) int {
	nStr := strconv.Itoa(n)
	return int(nStr[len(nStr)-1] - '0')
}

var cache = make(map[int]int)

func iterate(number, iterations int) (int, []int) {
	costs := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		if result, ok := cache[number]; ok {
			number = result
			costs[i] = cost(number)
			continue
		}
		start := number
		number = prune(mix(number, number*64))
		number = prune(mix(number, number/32))
		number = prune(mix(number, number*2048))
		cache[start] = number
		costs[i] = cost(number)
	}
	return number, costs
}

func makeKey(is []int) string {
	var str string
	for _, i := range is {
		str += strconv.Itoa(i) + ","
	}
	return str
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	costsPerBuyer := make(map[int][]int, len(lines))
	var sum int
	for _, line := range lines {
		secret, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		result, costs := iterate(secret, 2000)
		sum += result
		costsPerBuyer[secret] = costs
	}

	perKey := make(map[string]int)
	for secret, costs := range costsPerBuyer {
		uniqueKeys := make(map[string]struct{})
		diff := make([]int, len(costs))
		diff[0] = costs[0] - cost(secret)
		for i := 1; i < len(costs); i++ {
			diff[i] = costs[i] - costs[i-1]
		}

		for i := 0; i < len(costs)-3; i++ {
			key := makeKey(diff[i : i+4])
			if _, ok := uniqueKeys[key]; ok {
				// We have already seen this run before, so we can't check it again
				// since monkeys buy the moment they see the run.
				continue
			}
			uniqueKeys[key] = struct{}{}
			perKey[key] += costs[i+3]
		}
	}

	var best int
	for _, sum := range perKey {
		if sum > best {
			best = sum
		}
	}

	return best, nil
}
