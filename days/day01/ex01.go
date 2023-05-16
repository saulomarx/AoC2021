package day01

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
	messure := convertInputToInt(inputs)

	var howDeepWeGo = 0
	for i, el := range messure {
		if i != 0 {
			if el > messure[i-1] {
				howDeepWeGo++
			}
		}
	}
	fmt.Println(howDeepWeGo)

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day01/input.txt")
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

func convertInputToInt(input []string) []int {
	var response []int
	for _, el := range input {
		value, err := strconv.Atoi(el)

		if err == nil {
			response = append(response, value)
		}

	}
	return response
}
