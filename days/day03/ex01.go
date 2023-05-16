package day03

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
	gammaString, epsilonString := getGammaEpsilon(inputs)

	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)

	fmt.Println(gamma, epsilon)
	fmt.Println(gamma * epsilon)

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day03/input.txt")
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

func getGammaEpsilon(input []string) (gamma, epsilon string) {
	positionMap := make(map[int]int)
	inputHalfLength := len(input) / 2

	for _, el := range input {
		chars := strings.Split(el, "")
		for j, char := range chars {
			value, err := strconv.Atoi(char)
			if err == nil {
				positionMap[j] = positionMap[j] + value
			}

		}

	}

	keys := make([]int, 0, len(positionMap))

	for k := range positionMap {
		keys = append(keys, k)

	}

	for _, key := range keys {
		if positionMap[key] > inputHalfLength {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}

	}

	return
}
