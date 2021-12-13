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
	lavaTubes := make([][]int, 0)

	for _, input := range inputs {
		line := convertStringLineToNumbers(input)

		if len(line) > 0 {
			lavaTubes = append(lavaTubes, line)

		}
	}

	fmt.Println(lavaTubes)
	fmt.Println(lavaTubes[0][0])

	botton := findLowerPoints(lavaTubes)

	var sum int

	for _, value := range botton {
		sum += value + 1
	}
	fmt.Println(sum)
}

func lowerThenFourPoints(p, p1, p2, p3, p4 int) bool {
	return p < p1 && p < p2 && p < p3 && p < p4
}

func lowerThenThreePoints(p, p1, p2, p3 int) bool {
	return p < p1 && p < p2 && p < p3
}

func lowerThenTwoPoints(p, p1, p2 int) bool {
	return p < p1 && p < p2
}

func findLowerPoints(lavaTubes [][]int) []int {
	columnSize := len(lavaTubes)
	resp := make([]int, 0)

	for i, line := range lavaTubes {
		lineSize := len(line)
		for j := 0; j < lineSize; j++ {
			isTop := i == 0
			isBotton := i == columnSize-1
			isLeft := j == 0
			isRight := j == lineSize-1
			fmt.Println(lavaTubes[i][j], i, j)
			fmt.Println("-------------")
			if isTop && isLeft {
				// connor topLeft

				if lowerThenTwoPoints(lavaTubes[i][j], lavaTubes[i+1][j], lavaTubes[i][j+1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else if isTop && isRight {
				// connor topRight

				if lowerThenTwoPoints(lavaTubes[i][j], lavaTubes[i+1][j], lavaTubes[i][j-1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else if isBotton && isLeft {
				// bottonLeft
				if lowerThenTwoPoints(lavaTubes[i][j], lavaTubes[i][j+1], lavaTubes[i-1][j]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else if isBotton && isRight {
				// bottonRight
				if lowerThenTwoPoints(lavaTubes[i][j], lavaTubes[i][j-1], lavaTubes[i-1][j]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else if isTop {
				if lowerThenThreePoints(lavaTubes[i][j], lavaTubes[i][j-1], lavaTubes[i+1][j], lavaTubes[i][j+1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}

			} else if isBotton {
				if lowerThenThreePoints(lavaTubes[i][j], lavaTubes[i-1][j], lavaTubes[i][j+1], lavaTubes[i][j-1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else if isLeft {
				if lowerThenThreePoints(lavaTubes[i][j], lavaTubes[i+1][j], lavaTubes[i-1][j], lavaTubes[i][j+1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else if isRight {
				if lowerThenThreePoints(lavaTubes[i][j], lavaTubes[i-1][j], lavaTubes[i+1][j], lavaTubes[i][j-1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}
			} else {
				if lowerThenFourPoints(lavaTubes[i][j], lavaTubes[i-1][j], lavaTubes[i+1][j], lavaTubes[i][j-1], lavaTubes[i][j+1]) {
					fmt.Println(lavaTubes[i][j])
					resp = append(resp, lavaTubes[i][j])
				}

			}

		}
	}
	return resp
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
