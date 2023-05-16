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
	horizontal, deep := calcCoordinates(inputs)

	fmt.Println(horizontal * deep)

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

func calcCoordinates(input []string) (horizontal, deep int) {
	aim := 0

	for _, el := range input {
		slicedValue := strings.Split(el, " ")
		key := slicedValue[0]

		value, err := strconv.Atoi(slicedValue[1])

		if err == nil {
			if key == "down" {
				aim = aim + value

			} else if key == "up" {
				aim = aim - value

			} else if key == "forward" {
				fmt.Println("ain ->", aim, "deep ->", deep, "+", aim, "*", value, "=", deep+aim*value, "|", aim, "*", value, "=", aim*value)

				horizontal = horizontal + value
				deep = deep + aim*value
			}
		}

	}
	fmt.Println(horizontal, deep)
	return
}
