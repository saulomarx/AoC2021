package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputs := readInput()
	_, out := getInputsAndOutputs(inputs)

	var sum int

	for _, el := range out {
		el = strings.TrimSpace(el)
		// fmt.Println(el)
		size := len(el)
		// 1 4 7 8
		if size == 2 || size == 3 || size == 4 || size == 7 {
			sum++
		}

	}

	fmt.Println(sum)

}

func splitSignals(input string) []string {
	slicedInput := strings.Split(input, " ")

	for i, value := range slicedInput {
		slicedInput[i] = strings.TrimSpace(value)
	}

	return slicedInput
}

func getInputsAndOutputs(entries []string) (inputsSignal, outputsSignal []string) {
	for _, entries := range entries {
		slicedInput := strings.Split(entries, "|")
		inputLine, outputLine := slicedInput[0], slicedInput[1]
		inputsSignal = append(inputsSignal, splitSignals(inputLine)...)
		outputsSignal = append(outputsSignal, splitSignals(outputLine)...)
	}

	return
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day08/input.txt")
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
