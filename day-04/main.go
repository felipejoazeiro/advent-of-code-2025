package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	data, err := getDate()
	if err != nil {
		fmt.Println("Erro ao obter dados:", err)
		return
	}
	fmt.Println("Processando:", firstPart(data))
}

func firstPart(data []string) int {
	value := 0

	m:=make([][]rune, len(data))

	for i, s := range data {
		m[i] = []rune(s)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			valid := 0

			if m[i][j] == '.' {
				continue
			}

			if i == 0 {
				if j == 0 {
					if m[i+1][j] == '@' {
						valid++
					}
					if m[i][j+1] == '@' {
						valid++
					}
					if m[i+1][j+1] == '@' {
						valid++
					}
				}else if j == len(m[i])-1 {
					if m[i+1][j] == '@' {
						valid++
					}
					if m[i][j-1] == '@' {
						valid++
					}
					if m[i+1][j-1] == '@' {
						valid++
					}
				}else {
					if m[i][j-1] == '@' {
						valid++
					}
					if m[i][j+1] == '@' {
						valid++
					}
					if m[i+1][j] == '@' {
						valid++
					}
					if m[i+1][j-1] == '@' {
						valid++
					}
					if m[i+1][j+1] == '@' {
						valid++
					}
				}
			}else if i == len(m)-1 {
				if j == 0 {
					if m[i-1][j] == '@' {
						valid++
					}
					if m[i][j+1] == '@' {
						valid++
					}
					if m[i-1][j+1] == '@' {
						valid++
					}
				}else if j == len(m[i])-1 {
					if m[i-1][j] == '@' {
						valid++
					}
					if m[i][j-1] == '@' {
						valid++
					}
					if m[i-1][j-1] == '@' {
						valid++
					}
				}else {
					if m[i][j-1] == '@' {
						valid++
					}
					if m[i][j+1] == '@' {
						valid++
					}
					if m[i-1][j] == '@' {
						valid++
					}
					if m[i-1][j-1] == '@' {
						valid++
					}
					if m[i-1][j+1] == '@' {
						valid++
					}
				}
			}else {
				if j == 0 {
					if m[i-1][j] == '@' {
						valid++
					}
					if m[i+1][j] == '@' {
						valid++
					}
					if m[i][j+1] == '@' {
						valid++
					}
					if m[i-1][j+1] == '@' {
						valid++
					}
					if m[i+1][j+1] == '@' {
						valid++
					}
				}else if j == len(m[i])-1 {
					if m[i-1][j] == '@' {
						valid++
					}
					if m[i+1][j] == '@' {
						valid++
					}
					if m[i][j-1] == '@' {
						valid++
					}
					if m[i-1][j-1] == '@' {
						valid++
					}
					if m[i+1][j-1] == '@' {
						valid++
					}
				}else {
					if m[i-1][j] == '@' {
						valid++
					}
					if m[i+1][j] == '@' {
						valid++
					}
					if m[i][j-1] == '@' {
						valid++
					}
					if m[i][j+1] == '@' {
						valid++
					}
					if m[i-1][j-1] == '@' {
						valid++
					}
					if m[i-1][j+1] == '@' {
						valid++
					}
					if m[i+1][j-1] == '@' {
						valid++
					}
					if m[i+1][j+1] == '@' {
						valid++
					}
				}
			}


			if valid <= 3 {
				value++
			}
		}
	}

	return value
}

func 

func getDate() ([]string, error) {
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
