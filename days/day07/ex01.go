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
	crabs := convertStringLineToNumbers(inputs[0])

	crabMap := createsCrabMap(crabs)

	crabPositions := getCrabPositions(crabMap)

	minFuel := getMinFuel(crabMap, crabPositions)
	fmt.Println("Min fuel usage", minFuel)

}

func getDeltaPosition(a, b int) int {
	sum := a - b
	if sum < 0 {
		sum = -sum
	}
	return sum
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinFuel(crabMap map[int]int, positions []int) int {
	var fuel int
	for _, testPosition := range positions {
		auxFuel := 0
		for _, position := range positions {

			deltaDistance := getDeltaPosition(testPosition, position) * crabMap[position]
			auxFuel += deltaDistance
		}
		if fuel == 0 {
			fuel = auxFuel
		}

		fuel = getMin(fuel, auxFuel)

	}
	return fuel
}

func getCrabPositions(crabMap map[int]int) []int {
	positions := make([]int, 0, len(crabMap))
	for k := range crabMap {
		positions = append(positions, k)
	}

	sort.Ints(positions)
	return positions
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day07/input.txt")
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

func createsCrabMap(crabs []int) map[int]int {
	crabMap := make(map[int]int)

	for _, crab := range crabs {
		crabMap[crab] = crabMap[crab] + 1

	}
	return crabMap
}
