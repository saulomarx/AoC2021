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

	fuelMap := generateFuelUsageMap(crabPositions[len(crabPositions)-1])

	minFuel := getMinFuel(crabMap, fuelMap, crabPositions)
	fmt.Println("Min fuel usage", minFuel)

}

func generateFuelUsageMap(max int) map[int]int {
	fuelUsageMap := make(map[int]int)
	fuelUsageMap[0] = 0
	for i := 1; i <= max; i++ {
		fuelUsageMap[i] = fuelUsageMap[i-1] + i
	}
	return fuelUsageMap
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

func getMinFuel(crabMap, fuelMap map[int]int, positions []int) int {
	var fuel int
	first, last := positions[0], positions[len(positions)-1]
	for testPosition := first; testPosition <= last; testPosition++ {
		auxFuel := 0
		for _, position := range positions {

			deltaDistance := getDeltaPosition(testPosition, position)
			totalFuel := fuelMap[deltaDistance] * crabMap[position]
			auxFuel += totalFuel
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
