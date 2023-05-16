package day02

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
	coordinatesMap := convertInputToMap(inputs)

	absoluteDown := coordinatesMap["down"] - coordinatesMap["up"]
	fmt.Println(coordinatesMap["forward"], absoluteDown)
	fmt.Println(absoluteDown * coordinatesMap["forward"])

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day02/input.txt")
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

func convertInputToMap(input []string) map[string]int {
	response := make(map[string]int)
	for _, el := range input {
		slicedValue := strings.Split(el, " ")
		key := slicedValue[0]

		value, err := strconv.Atoi(slicedValue[1])

		if err == nil {
			response[key] = response[key] + value
		}

	}
	return response
}
