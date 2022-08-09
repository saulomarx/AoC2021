package ex

import (
	"fmt"
)

func Ex02(raw []string) {
	fmt.Println("EX02")
	insert, instruction := GetInsertAndInstructions(raw)
	instructionMap, instructionCountMap := GetInstructionMapAndCount(instruction)

	initialCount(insert, instructionCountMap)
	in := insert
	fmt.Println("Rec aqui-------")
	for i := 0; i < 40; i++ {
		fmt.Printf("------- Cycle %v\n", i+1)
		in = recPoly(in, instructionMap, instructionCountMap)
	}

	fmt.Println(len(in))
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
