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
	fishes := convertStringLineToNumbers(inputs[0])

	fmt.Println(fishes)
	fishMap := createsFishMap(fishes)

	for i := 8; i >= 0; i-- {
		fmt.Println(i, fishMap[i])
	}
	fishCycles(fishMap, 80)
}

func fishCycles(fishMap map[int]int, days int) {

	for k := 0; k < days; k++ {
		auxFishMap := make(map[int]int)
		for i := 8; i >= 0; i-- {
			if i == 0 {
				auxFishMap[8] = auxFishMap[8] + fishMap[0]
				auxFishMap[6] = auxFishMap[6] + fishMap[0]
			} else {
				auxFishMap[i-1] = fishMap[i]
			}

		}

		fishMap = auxFishMap
	}
	var totalFishes int
	for i := 8; i >= 0; i-- {
		totalFishes += fishMap[i]
	}
	fmt.Println("Total:", totalFishes)
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day06/input.txt")
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
	stringNumbers := strings.Split(input, ",")

	for _, stringNumber := range stringNumbers {
		value, _ := strconv.Atoi(stringNumber)
		numbersList = append(numbersList, value)

	}
	return numbersList
}

func createsFishMap(fishes []int) map[int]int {
	fishMap := make(map[int]int)

	for i := 0; i < 9; i++ {
		fishMap[i] = 0

	}

	for _, fish := range fishes {
		fishMap[fish] = fishMap[fish] + 1

	}

	return fishMap
}
