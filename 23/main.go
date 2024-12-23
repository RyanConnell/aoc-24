package main

import (
	"fmt"
	"sort"
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

type Node struct {
	name        string
	connections map[string]*Node
}

func (n *Node) connectTo(target *Node) {
	n.connections[target.name] = target
	target.connections[n.name] = n
}

func groupName(strs ...string) string {
	sort.Strings(strs)
	return strings.Join(strs, ",")
}

func (n *Node) withinRange(target string, distance int) (string, bool) {
	if distance == 0 {
		return "", false
	}
	for _, peer := range n.connections {
		if peer.name == target && distance == 1 {
			return n.name, true
		}
		if peer.name == target {
			continue
		}
		if middle, ok := peer.withinRange(target, distance-1); ok {
			return middle, true
		}
	}
	return "", false
}

func (n *Node) groups() []string {
	var groups []string
	for _, peer := range n.connections {
		middle, ok := peer.withinRange(n.name, 2)
		if !ok {
			continue
		}
		if peer.name[0] == 't' || middle[0] == 't' || n.name[0] == 't' {
			groups = append(groups, groupName(n.name, peer.name, middle))
		}
	}
	return groups
}

func solve(lines []string) (int, error) {
	nodePool := make(map[string]*Node)
	getNode := func(target string) *Node {
		if val, ok := nodePool[target]; ok {
			return val
		}
		nodePool[target] = &Node{name: target, connections: make(map[string]*Node)}
		return nodePool[target]
	}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		getNode(parts[0]).connectTo(getNode(parts[1]))
	}

	uniqueGroups := make(map[string]int)
	for _, node := range nodePool {
		fmt.Println(node)
		if len(node.connections) < 2 {
			continue
		}
		groups := node.groups()
		if len(groups) != 0 {
			fmt.Println(node.name, groups)
		}
		for _, group := range groups {
			uniqueGroups[group]++
		}
	}
	fmt.Println(uniqueGroups)
	for key, val := range uniqueGroups {
		fmt.Println(val, key)
	}
	return len(uniqueGroups), nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
