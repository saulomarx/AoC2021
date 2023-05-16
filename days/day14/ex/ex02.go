package ex

import (
	"fmt"
)

func Ex02(raw []string) {
	fmt.Println("EX02")
	insert, instruction := GetInsertAndInstructions(raw)
	instructionMap, instructionCountMap := GetInstructionMapAndCount(instruction)

	in := insert
	fmt.Println("Rec aqui-------")
	for i := 0; i < 40; i++ {
		fmt.Printf("------- Cycle %v\n", i+1)
		in = memoryRec(in, instructionMap)
	}
	fmt.Println(fmt.Println("Rec fim-------"))

	//initialCount(in, instructionCountMap)
	max, min := GetMaxMinFromCount(instructionCountMap)

	fmt.Printf("\nMax %v - Min %v = %v\n", max, min, max-min)
	fmt.Println(instructionCountMap)
}

func initialCount(instruction string, count map[string]int) {
	for _, l := range instruction {
		el := string(l)
		count[el]++
	}

}

func singlePolymerization(input string, instructionMap map[string]string, count map[string]int) string {
	if len(input) < 2 {
		fmt.Println("DEU RUIM")
	}
	e := instructionMap[input]
	count[e]++
	return string(input[0]) + e
}

func sPolymerization(input string, instructionMap map[string]string) string {
	if len(input) < 2 {
		fmt.Println("DEU RUIM")
	}
	e := instructionMap[input]
	return string(input[0]) + e
}

func recPoly(input string, instructionMap map[string]string, count map[string]int) string {
	inputLen := len(input)

	if inputLen == 2 {
		return singlePolymerization(input, instructionMap, count) + string(input[inputLen-1])
	}
	if inputLen == 3 {
		return singlePolymerization(input[:2], instructionMap, count) + singlePolymerization(input[1:], instructionMap, count) + string(input[inputLen-1])
	}

	halfLen := inputLen / 2

	middle := input[halfLen-1 : halfLen+1]
	e := instructionMap[middle]
	count[e]++
	return recPoly(input[:halfLen], instructionMap, count) + e + recPoly(input[halfLen:], instructionMap, count)
}

func memoryRec(input string, instructions map[string]string) string {
	if value, ok := instructions[input]; ok && len(input) > 2 {
		return value
	}

	inputLen := len(input)

	if inputLen == 2 {
		newPolymer := sPolymerization(input, instructions) + string(input[inputLen-1])
		//instructions[input] = newPolymer
		return newPolymer

	}
	if inputLen == 3 {
		newPolymer := sPolymerization(input[:2], instructions) + sPolymerization(input[1:], instructions) + string(input[inputLen-1])
		//instructions[input] = newPolymer
		return newPolymer
	}

	halfLen := inputLen / 2

	middle := input[halfLen-1 : halfLen+1]
	e := instructions[middle]
	newPolymer := memoryRec(input[:halfLen], instructions) + e + memoryRec(input[halfLen:], instructions)
	instructions[input] = newPolymer
	return newPolymer

}
