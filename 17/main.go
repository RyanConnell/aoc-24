package main

import (
	"fmt"
	"math"
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
	fmt.Printf("Part 1 Result: %s\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %s\n", solutionPart2)
}

/// Part 1 \\\

const (
	ADV = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

var codeRepr = map[int]string{
	0: "ADV",
	1: "BXL",
	2: "BST",
	3: "JNZ",
	4: "BXC",
	5: "OUT",
	6: "BDV",
	7: "CDV",
}

type Memory struct {
	Instructions []int // Incoming instruction set
	Location     int   // Location in memory we are reading from.
	Registers    []int // Values in each register
	Output       []int
}

func (m *Memory) run() string {
	for m.Location < len(m.Instructions) {
		opcode, operand := m.Instructions[m.Location], m.Instructions[m.Location+1]
		m.execute(opcode, operand)
		m.Location += 2
	}

	var out string
	for _, o := range m.Output {
		out += fmt.Sprintf("%d,", o)
	}
	//fmt.Println(m.Registers)
	if out != "" {
		return out[:len(out)-1]
	}
	return ""
}

func (m *Memory) execute(opcode, operand int) {
	//fmt.Println("execute:", codeRepr[opcode], operand, m.Registers)
	switch opcode {
	case ADV, BDV, CDV:
		m.dv(opcode, operand)
	case BXL:
		m.bxl(operand)
	case BST:
		m.bst(operand)
	case JNZ:
		m.jnz(operand)
	case BXC:
		m.bxc(operand)
	case OUT:
		m.out(operand)
	default:
		panic("unknown opcode")
	}
}

func (m *Memory) dv(opcode, operand int) {
	value := int(float64(m.Registers[0]) / math.Pow(2, float64(m.combo(operand))))

	// Write the result into the corresponding register.
	target := 0
	if opcode == BDV {
		target = 1
	} else if opcode == CDV {
		target = 2
	}
	m.Registers[target] = value
}

func (m *Memory) bxl(operand int) {
	value := m.Registers[1] ^ operand
	m.Registers[1] = value
}

func (m *Memory) bst(operand int) {
	value := m.combo(operand) % 8
	m.Registers[1] = value
}

func (m *Memory) jnz(operand int) {
	if m.Registers[0] == 0 {
		return
	}
	m.Location = operand - 2 // -2 because we aren't supposed to jump after this.
}

func (m *Memory) bxc(_ int) {
	value := m.Registers[1] ^ m.Registers[2]
	m.Registers[1] = value
}

func (m *Memory) out(operand int) {
	m.Output = append(m.Output, m.combo(operand)%8)
}

func (m *Memory) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4, 5, 6:
		return m.Registers[operand-4]
	}
	panic("We shouldn't be here")
}

func solve(lines []string) (string, error) {
	memory := &Memory{
		Registers: make([]int, 0),
	}
	for _, line := range lines {
		val := strings.Split(line, ": ")
		if strings.HasPrefix(val[0], "Register") {
			registerValue, err := strconv.Atoi(val[1])
			if err != nil {
				return "", err
			}
			memory.Registers = append(memory.Registers, registerValue)
		} else if strings.HasPrefix(val[0], "Program") {
			instructionStrs := strings.Split(val[1], ",")
			instructions := make([]int, len(instructionStrs))
			for i, s := range instructionStrs {
				instructions[i] = int(s[0] - '0')
			}
			memory.Instructions = instructions
		}
	}

	//fmt.Printf("Memory: %+v\n", memory)
	return memory.run(), nil
}

/// Part 2 \\\

func solvePart2(lines []string) (string, error) {
	return "", nil
}
