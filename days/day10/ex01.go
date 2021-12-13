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

	tokenCloseMap := createTokenCloseMap()
	var sum int
	for k := range inputs {
		result := analizeString(inputs[k], tokenCloseMap)
		sum += setPoints(result)
	}

	fmt.Println(sum)

}

func setPoints(s string) int {
	// 	): 3 points.
	// ]: 57 points.
	// }: 1197 points.
	// >: 25137 points.
	switch s {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}

func analizeString(input string, tokenCloseMap map[string]string) string {
	stack := make([]string, 0)
	for _, c := range input {

		value := string(c)
		if tokenCloseMap[value] == "" {
			stack = append(stack, value)
		} else {
			lastPositionIdx := len(stack) - 1
			if stack[lastPositionIdx] == tokenCloseMap[value] {
				stack = pop(stack)
			} else {
				return value
			}
		}
	}

	return ""
}

func createTokenCloseMap() map[string]string {
	tokenCloseMap := make(map[string]string)
	tokenCloseMap["}"] = "{"
	tokenCloseMap["]"] = "["
	tokenCloseMap[">"] = "<"
	tokenCloseMap[")"] = "("

	return tokenCloseMap
}

func pop(stack []string) []string {
	size := len(stack)
	stack = stack[:size-1]
	return stack
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day10/input.txt")
	if err != nil {
		fmt.Println("Deu ruim")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		input = append(input, line)
	}

	return input
}
