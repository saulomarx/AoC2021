package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"AoC2021/days/day14/ex"
)

const inputPath = "./inputs/day14/inputTest01.txt"

func main() {
	raw := readInput(inputPath)

	//ex.Ex01(raw)
	ex.Ex02(raw)
}

func readInput(in string) []string {
	var input []string
	file, err := os.Open(in)
	if err != nil {
		fmt.Println("Deu ruim")
	}
	leitor := bufio.NewReader(file)

	for {
		line, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		input = append(input, line)
	}

	return input
}
