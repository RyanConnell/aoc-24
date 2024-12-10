package parser

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadFile(fileName string) ([]string, bool) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, false
	}
	defer file.Close()
	return readLines(file), true
}

func MustReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to read given file")
	}
	defer file.Close()
	return readLines(file)
}

func readLines(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
