package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	dumboMtx := make([][]int, 0)
	var flashes int

	for _, input := range inputs {
		line := convertStringLineToNumbers(input)

		if len(line) > 0 {
			dumboMtx = append(dumboMtx, line)

		}
	}

	steps := 100
	for k := 0; k < steps; k++ {
		var stepFlashes int
		dumboMtx = plusOne(dumboMtx)
		dumboMtx, stepFlashes = flashCycle(dumboMtx)
		flashes += stepFlashes
	}
	fmt.Println(flashes)

}

func flashCycle(matrix [][]int) ([][]int, int) {
	var flashes int
	changed := true
	columnSize := len(matrix[0]) - 1
	lineSize := len(matrix) - 1

	for changed {
		changed = false
		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] == 10 {
					changed = true
					matrix[i][j]++

					// left
					if i > 0 && matrix[i-1][j] < 10 {
						matrix[i-1][j]++
					}
					// right
					if i < columnSize && matrix[i+1][j] < 10 {
						matrix[i+1][j]++
					}
					//top
					if j > 0 && matrix[i][j-1] < 10 {
						matrix[i][j-1]++
					}
					// down
					if j < lineSize && matrix[i][j+1] < 10 {
						matrix[i][j+1]++
					}
					// topLeft
					if i > 0 && j > 0 && matrix[i-1][j-1] < 10 {
						matrix[i-1][j-1]++
					}
					// topRight
					if i > 0 && j < lineSize && matrix[i-1][j+1] < 10 {
						matrix[i-1][j+1]++
					}
					// bottonLeft
					if i < columnSize && j > 0 && matrix[i+1][j-1] < 10 {
						matrix[i+1][j-1]++
					}
					// bottonRight
					if i < columnSize && j < lineSize && matrix[i+1][j+1] < 10 {
						matrix[i+1][j+1]++
					}

				}
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > 9 {
				matrix[i][j] = 0
				flashes++
			}
		}
	}

	return matrix, flashes
}

func printMtx(matrix [][]int) {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func plusOne(matrix [][]int) [][]int {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j]++
		}
	}
	return matrix

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day11/input.txt")
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
