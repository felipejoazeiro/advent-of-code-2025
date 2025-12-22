package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func countZeroClicks(pos int, dir byte, steps int, n int) int {
	if steps <= 0 {
		return 0
	}

	var t0 int
	switch dir {
	case 'R':
		t0 = (n - pos) % n
		if t0 == 0 {
			t0 = n
		}
	case 'L':
		t0 = pos % n
		if t0 == 0 {
			t0 = n
		}
	default:
		return 0
	}

	if steps < t0 {
		return 0
	}
	return 1 + (steps-t0)/n
}

func move(pos int, dir byte, steps int, n int) int {
	switch dir {
	case 'R':
		return (pos + steps) % n
	case 'L':
		return ((pos-steps)%n + n) % n
	default:
		return pos
	}
}

func main() {
	data, err := getData()
	if err != nil {
		log.Fatal(err)
	}

	const n = 100
	pos := 50
	password := 0

	for _, line := range data {
		dir := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("erro ao converter passos %q: %v", line, err)
		}

		password += countZeroClicks(pos, dir, steps, n)

		pos = move(pos, dir, steps, n)
	}

	fmt.Println("Password (method 0x434C49434B):", password)
}

func getData() ([]string, error) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("não foi possível obter o caminho do arquivo")
	}
	dir := filepath.Dir(thisFile)
	path := filepath.Join(dir, "input.txt")

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := strings.Split(string(b), "\n")
	tokens := make([]string, 0, len(raw))
	for _, t := range raw {
		s := strings.TrimSpace(t)
		if s == "" {
			continue
		}
		tokens = append(tokens, s)
	}
	return tokens, nil
}
