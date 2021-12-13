package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	lavaTubes := make([][]int, 0)

	for _, input := range inputs {
		line := convertStringLineToNumbers(input)

		if len(line) > 0 {
			lavaTubes = append(lavaTubes, line)

		}
	}

	columnSize := len(lavaTubes[0])
	lineSize := len(lavaTubes)

	areaSizes := make([]int, 0)

	for i := range lavaTubes {
		for j := range lavaTubes[i] {
			if lavaTubes[i][j] >= 0 || lavaTubes[i][j] < 9 {
				area := fillArea(lavaTubes, i, j, columnSize, lineSize)
				if area > 0 {
					areaSizes = append(areaSizes, area)
				}
			}
		}
	}

	sort.Ints(areaSizes)
	numberOfAreas := len(areaSizes)
	v1, v2, v3 := areaSizes[numberOfAreas-1], areaSizes[numberOfAreas-2], areaSizes[numberOfAreas-3]
	fmt.Println(v1, "*", v2, "*", v3, "=", v1*v2*v3)

}

func fillArea(matrix [][]int, i, j, columnSize, lineSize int) int {
	var sum int
	if i < 0 || i >= lineSize || j < 0 || j >= columnSize {
		return sum
	}
	value := matrix[i][j]
	if value > 8 || value < 0 {
		return sum
	}

	matrix[i][j] = -1
	sum++

	sum += fillArea(matrix, i+1, j, columnSize, lineSize)
	sum += fillArea(matrix, i-1, j, columnSize, lineSize)
	sum += fillArea(matrix, i, j+1, columnSize, lineSize)
	sum += fillArea(matrix, i, j-1, columnSize, lineSize)

	return sum

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day09/input.txt")
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

func convertStringLineToNumbers(input string) []int {
	var numbersList []int

	for _, stringNumber := range input {
		value, _ := strconv.Atoi(string(stringNumber))
		numbersList = append(numbersList, value)

	}
	return numbersList
}
