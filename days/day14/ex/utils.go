package ex

import (
	"strings"
)

func GetInsertAndInstructions(l []string) (insert string, instructions []string) {
	insert = l[0]
	instructions = l[2:]
	return
}

func GetInstructionMapAndCount(instructions []string) (map[string]string, map[string]int) {
	instructionMap := make(map[string]string)
	instructionCountMap := make(map[string]int)

	for _, el := range instructions {
		s := strings.Split(el, " -> ")
		instructionMap[s[0]] = s[1]
		instructionCountMap[s[1]] = 0
	}
	return instructionMap, instructionCountMap
}

func GetMaxMinFromCount(count map[string]int) (max, min int) {
	for _, el := range count {
		if max < el {
			max = el
		}
		if min > el || min == 0 {
			min = el
		}
	}
	return
}
