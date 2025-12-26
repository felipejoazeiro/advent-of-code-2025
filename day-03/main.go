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

func firstPart(data []string) int {
	value := 0

	for _, line := range data {
		raw := strings.Split(line, "")
		fmt.Println("Valor integro:", raw)

		hightestValue := 0
		secondHightestValue := 0
		hightestValueIndex := 0
		for i := 0; i < len(raw)-1; i++ {
			rest, _ := strconv.Atoi(raw[i])

			if rest > hightestValue {
				hightestValue = rest
				hightestValueIndex = i
			}
		}
		for i := hightestValueIndex + 1; i < len(raw); i++ {
			rest, _ := strconv.Atoi(raw[i])
			if rest > secondHightestValue {
				secondHightestValue = rest
			}
		}
		valueStr := "" + strconv.Itoa(hightestValue) + strconv.Itoa(secondHightestValue)
		intValue, _ := strconv.Atoi(valueStr)
		value += intValue
	}
	return value
}

func secondPart(data []string) int {
	value := 0

	for _, line := range data {
		raw := strings.Split(line, "")
		fmt.Println("Valor integro:", raw)
		data := []int{}
		dataValue := 0
		indexValue := 0

		for i := 12; i > 0; i-- {
			dataValue = 0
			for j := indexValue; j < len(raw)-i+1; j++ {
				rest, _ := strconv.Atoi(raw[j])
				if rest > dataValue {
					dataValue = rest
					indexValue = j + 1
				}
			}
			data = append(data, dataValue)

		}

		valueStr := ""

		for _, v := range data {
			valueStr += strconv.Itoa(v)
		}
		intValue, _ := strconv.Atoi(valueStr)
		fmt.Println("Valor parcial:", valueStr)
		value += intValue
	}
	return value
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
