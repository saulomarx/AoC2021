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
	in, out := getInputsAndOutputs(inputs)
	var sum int

	for idx, signal := range in {
		map7 := return7(signal)
		map1 := return1(signal)
		map8 := return8(signal)
		map4 := return4(signal)
		find5digits := find5Digits(signal)
		find6digits := find6Digits(signal)

		// discovery3
		var map3 string
		map3, find5digits = discovery3(map1, find5digits)

		// discovery 5
		var map5 string
		map5, find5digits = discovery5(map1, map4, find5digits)

		// discovery 2
		map2 := find5digits[0]

		// discovery 9
		var map9 string
		map9, find6digits = discovery9(map2, map3, find6digits)

		// discovery 6
		var map6 string
		map6, find6digits = discovery6(map9, map5, find6digits)

		// discovery 0
		map0 := find6digits[0]

		translateMap := make(map[string]string)
		translateMap[sortString(strings.TrimSpace(map0))] = "0"
		translateMap[sortString(strings.TrimSpace(map1))] = "1"
		translateMap[sortString(strings.TrimSpace(map2))] = "2"
		translateMap[sortString(strings.TrimSpace(map3))] = "3"
		translateMap[sortString(strings.TrimSpace(map4))] = "4"
		translateMap[sortString(strings.TrimSpace(map5))] = "5"
		translateMap[sortString(strings.TrimSpace(map6))] = "6"
		translateMap[sortString(strings.TrimSpace(map7))] = "7"
		translateMap[sortString(strings.TrimSpace(map8))] = "8"
		translateMap[sortString(strings.TrimSpace(map9))] = "9"

		var result []rune
		for _, el := range out[idx] {
			if len(strings.TrimSpace(el)) > 0 {
				value := translateMap[sortString(strings.TrimSpace(el))]
				result = append(result, []rune(value)...)

			}
		}
		resultString := string(result)
		fmt.Println(resultString)
		value, _ := strconv.Atoi(resultString)
		sum = sum + value
		fmt.Println("-------------------")

	}

	fmt.Println(sum)

}

func discovery6(map9, map5 string, find6Digits []string) (map6 string, remaning6Digits []string) {
	removes := removeRunesFromAnother(map9, map5)
	for _, el := range find6Digits {
		v1 := strings.IndexRune(el, removes[0])
		if v1 < 0 {
			map6 = el

		} else {
			remaning6Digits = append(remaning6Digits, el)
		}
	}
	return
}

func discovery9(map2, map3 string, find6Digits []string) (map9 string, remaning6Digits []string) {
	removes := removeRunesFromAnother(map2, map3)

	for _, el := range find6Digits {
		v1 := strings.IndexRune(el, removes[0])
		if v1 < 0 {
			map9 = el
		} else {
			remaning6Digits = append(remaning6Digits, el)
		}
	}
	return
}

func discovery5(map1, map4 string, find5digits []string) (map5 string, remaning5Digits []string) {
	removes := removeRunesFromAnother(map4, map1)
	for _, el := range find5digits {
		v1 := strings.IndexRune(el, removes[0])
		v2 := strings.IndexRune(el, removes[1])
		if v1 >= 0 && v2 >= 0 {
			map5 = el
		} else {
			remaning5Digits = append(remaning5Digits, el)
		}
	}

	return

}

func discovery3(map1 string, find5digits []string) (map3 string, remaning5Digits []string) {
	map1Runes := []rune(map1)

	for _, el := range find5digits {
		v1 := strings.IndexRune(el, map1Runes[0])
		v2 := strings.IndexRune(el, map1Runes[1])
		if v1 >= 0 && v2 >= 0 {
			map3 = el
		} else {
			remaning5Digits = append(remaning5Digits, el)
		}
	}
	return

}

func sortString(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func removeRunesFromAnother(base, removes string) []rune {
	result := []rune(base)
	for _, c := range removes {
		inStringIdx := strings.IndexRune(base, c)
		if inStringIdx >= 0 {
			result = append(result[:inStringIdx], result[inStringIdx+1:]...)
			base = string(result)
		}

	}
	return result
}

func return7(input []string) string {
	for _, el := range input {
		if len(el) == 3 {
			return el
		}
	}
	return ""
}

func return1(input []string) string {
	for _, el := range input {
		if len(el) == 2 {
			return el
		}
	}
	return ""
}

func return4(input []string) string {
	for _, el := range input {
		if len(el) == 4 {
			return el
		}
	}
	return ""
}

func return8(input []string) string {
	for _, el := range input {
		if len(el) == 7 {
			return el
		}
	}
	return ""
}

func find5Digits(input []string) []string {
	var with5Digits []string
	for _, el := range input {
		if len(el) == 5 {
			with5Digits = append(with5Digits, el)
		}
	}
	return with5Digits
}

func find6Digits(input []string) []string {
	var with5Digits []string
	for _, el := range input {
		if len(el) == 6 {
			with5Digits = append(with5Digits, el)
		}
	}
	return with5Digits
}

func splitSignals(input string) []string {
	slicedInput := strings.Split(input, " ")

	for i, value := range slicedInput {
		slicedInput[i] = strings.TrimSpace(value)
	}

	return slicedInput
}

func getInputsAndOutputs(entries []string) (inputsSignal, outputsSignal [][]string) {
	for _, entries := range entries {
		actualInput := make([]string, 10)
		actualOutput := make([]string, 10)

		slicedInput := strings.Split(entries, "|")
		inputLine, outputLine := slicedInput[0], slicedInput[1]
		actualInput = append(actualInput, splitSignals(inputLine)...)
		actualOutput = append(actualOutput, splitSignals(outputLine)...)
		inputsSignal = append(inputsSignal, actualInput)
		outputsSignal = append(outputsSignal, actualOutput)
	}

	return
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day08/inputTest.txt")
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
