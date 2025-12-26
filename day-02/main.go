package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	data, err := getData()
	if err != nil {
		fmt.Println("Erro ao obter dados:", err)
		return
	}

	/* fmt.Println("Processando:", firstPart(data)) */
	fmt.Println("Processando:", secondPart(data))
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

	raw := strings.Split(string(b), ",")
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

func firstPart(data []string) int {
	valuesRepeated := 0

	for _, res := range data {
		raw := strings.Split(res, "-")
		values := []int{}

		val0, err := strconv.Atoi(raw[0])
		if err != nil {
			fmt.Println("Erro ao converter:", err)
			continue
		}
		val1, err := strconv.Atoi(raw[1])
		if err != nil {
			fmt.Println("Erro ao converter:", err)
			continue
		}
		for i := val0; i <= val1; i++ {
			values = append(values, i)
			s := strconv.Itoa(i)

			mid := len(s) / 2

			leftStr := s[:mid]
			rightStr := s[mid:]

			if leftStr == rightStr {
				valuesRepeated += i
			}
		}
	}
	return valuesRepeated
}

// isMadeOfRepeatedBlock returns true if the entire decimal string s
// is composed of k>=2 repetitions of some substring.
func isMadeOfRepeatedBlock(s string) bool {
	n := len(s)
	// Try all block sizes that divide n and are at most n/2
	for l := 1; l*2 <= n; l++ {
		if n%l != 0 {
			continue
		}
		block := s[:l]
		ok := true
		for k := l; k < n; k += l {
			if s[k:k+l] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func secondPart(data []string) int64 {
	var sumInvalid int64 = 0

	for _, res := range data {
		raw := strings.Split(res, "-")
		if len(raw) != 2 {
			fmt.Println("Intervalo inválido:", res)
			continue
		}
		val0, err := strconv.Atoi(strings.TrimSpace(raw[0]))
		if err != nil {
			fmt.Println("Erro ao converter:", err)
			continue
		}
		val1, err := strconv.Atoi(strings.TrimSpace(raw[1]))
		if err != nil {
			fmt.Println("Erro ao converter:", err)
			continue
		}
		if val0 > val1 {
			val0, val1 = val1, val0
		}
		for i := val0; i <= val1; i++ {
			s := strconv.Itoa(i)
			if isMadeOfRepeatedBlock(s) {
				sumInvalid += int64(i)
			}
		}
	}
	return sumInvalid
}
