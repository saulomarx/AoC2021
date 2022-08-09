package ex

import (
	"fmt"
)

func Ex01(raw []string) {
	insert, instruction := GetInsertAndInstructions(raw)
	instructionMap, instructionCountMap := GetInstructionMapAndCount(instruction)

	fmt.Println(insert)
	PolymerizationSteps(insert, 20, instructionMap, instructionCountMap)

	max, min := GetMaxMinFromCount(instructionCountMap)

	fmt.Printf("\nMax %v - Min %v = %v\n", max, min, max-min)
}

func PolymerizationSteps(input string, steps int, instructions map[string]string, count map[string]int) {
	polymer := input

	for _, el := range input {
		count[string(el)]++
	}

	for i := 0; i < steps; i++ {
		polymer = Polymerization(polymer, instructions, count)
	}

}

func Polymerization(input string, instructions map[string]string, count map[string]int) string {
	var newPolymer string
	inputLen := len(input)

	for idx := 1; idx < inputLen; idx++ {
		instruction := input[idx-1 : idx+1]
		e := instructions[instruction]
		count[e]++
		newPolymer = newPolymer + string(input[idx-1]) + e

	}
	fmt.Println(len(newPolymer))

	return newPolymer + string(input[inputLen-1])
}
