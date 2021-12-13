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
	missingClose := make([][]string, 0)
	scores := make([]int, 0)
	for k := range inputs {
		missing := analizeString(inputs[k], tokenCloseMap)
		if len(missing) > 0 {
			missingClose = append(missingClose, missing)
		}
	}

	for _, missingStack := range missingClose {
		var points int
		for i := len(missingStack) - 1; i >= 0; i-- {
			el := missingStack[i]
			points = points * 5
			value := setPoints(el)
			points += value

		}
		scores = append(scores, points)
	}
	middleIndex := len(scores) / 2
	fmt.Println(scores[middleIndex])

}

func setPoints(s string) int {
	// 	): 1 point.
	// ]: 2 points.
	// }: 3 points.
	// >: 4 points.
	switch s {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	default:
		return 0
	}
}

func analizeString(input string, tokenCloseMap map[string]string) []string {
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
				return make([]string, 0)
			}
		}
	}

	return stack
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
