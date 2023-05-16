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
	o2String := getO2(inputs)
	co2String := getCo2(inputs)

	o2, _ := strconv.ParseInt(o2String, 2, 64)
	co2, _ := strconv.ParseInt(co2String, 2, 64)

	fmt.Println(o2, co2)
	fmt.Println(o2 * co2)

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

func getO2(input []string) string {
	localInput := input
	bitPosition := 0
	for {
		sum := 0
		halfLength := len(localInput) / 2
		if len(localInput)%2 != 0 {
			halfLength++
		}
		for _, el := range localInput {
			char := el[bitPosition]
			if char == '1' {
				sum++
			}

		}
		var newInput []string
		// 1 is more significant
		if sum >= halfLength {
			for _, el := range localInput {
				if el[bitPosition] == '1' {
					newInput = append(newInput, el)
				}
			}

		} else {
			for _, el := range localInput {
				if el[bitPosition] == '0' {
					newInput = append(newInput, el)
				}
			}
		}

		localInput = newInput

		if len(localInput) == 1 {
			break
		}
		bitPosition++
	}

	return localInput[0]
}

func getCo2(input []string) string {
	localInput := input
	bitPosition := 0
	fmt.Println(localInput)
	for {
		sum := 0
		halfLength := len(localInput) / 2
		if len(localInput)%2 != 0 {
			halfLength++
		}
		for _, el := range localInput {
			char := el[bitPosition]
			if char == '1' {
				sum++
			}

		}
		var newInput []string
		// 1 is more significant
		if sum >= halfLength {
			for _, el := range localInput {
				if el[bitPosition] == '0' {
					newInput = append(newInput, el)
				}
			}

		} else {
			for _, el := range localInput {
				if el[bitPosition] == '1' {
					newInput = append(newInput, el)
				}
			}
		}

		localInput = newInput

		if len(localInput) == 1 {
			break
		}
		bitPosition++
	}

	return localInput[0]
}
