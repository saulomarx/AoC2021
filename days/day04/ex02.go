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

var line = [5]bool{false, false, false, false, false}
var matrix = [5][5]bool{line, line, line, line, line}

type bingoBoard struct {
	board       [5][5]int
	resultBoard [5][5]bool
	sum         int
	win         bool
}

func createBoard(board [5][5]int, sum int) *bingoBoard {
	b := bingoBoard{board: board, resultBoard: matrix, sum: sum, win: false}
	return &b
}

func main() {
	inputs := readInput()
	numbersString, boards := getNumbersAndBoards(inputs)

	boardMap, numberToMap := mapBoards(boards)
	numbers := convertStringLineToNumbers(numbersString)

	loseBingo(numbers, numberToMap, boardMap, len(boardMap))
}

func loseBingo(numbers []int, numberToMap map[int][]int, boardMap map[int]*bingoBoard, numberOfBoards int) {
	for round, number := range numbers {
		fmt.Println("Round -> ", round+1, "| number ->", number)
		boardsWithNumber := numberToMap[number]
		sort.Ints(boardsWithNumber)

		for _, boardIndex := range boardsWithNumber {
			if !boardMap[boardIndex].win {
				lineIndex, columnIndex := findNumberPosition(boardMap[boardIndex].board, number)
				columWin := false
				lineWin := false
				if lineIndex >= 0 && columnIndex >= 0 {

					boardMap[boardIndex].resultBoard[lineIndex][columnIndex] = true
					boardMap[boardIndex].sum -= number
					columWin = checkColumn(boardMap[boardIndex].resultBoard, columnIndex)
					lineWin = checkLine(boardMap[boardIndex].resultBoard, lineIndex)
				}

				if columWin || lineWin {
					fmt.Println("Winner! Board ", boardIndex)
					fmt.Println(boardMap[boardIndex].sum, number, boardMap[boardIndex].sum*number)
					boardMap[boardIndex].win = true
					fmt.Println(numberOfBoards)
					numberOfBoards = numberOfBoards - 1

					if numberOfBoards == 0 {
						return
					}
				}
			}

		}
	}

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day04/input.txt")
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

func checkColumn(matrix [5][5]bool, column int) bool {
	result := true
	for i := range matrix {
		result = result && matrix[i][column]
	}
	return result
}

func checkLine(matrix [5][5]bool, line int) bool {
	result := true
	for j := range matrix {
		result = result && matrix[line][j]
	}
	return result
}

func findNumberPosition(board [5][5]int, number int) (lineIndex, columIndex int) {
	for i, line := range board {
		for j, el := range line {
			if el == number {
				return i, j
			}
		}
	}
	return -1, -1
}

func getNumbersAndBoards(input []string) (numbers string, boards [][]string) {
	numbers = input[0]
	input = input[1:]

	var board []string

	for i, line := range input {
		mod := i % 6
		if mod > 0 {
			board = append(board, line)
		}
		if mod == 5 {
			boards = append(boards, board)
			board = nil
		}
	}

	return
}

func convertStringLineToNumbers(input string) []int {
	var numbersList []int
	stringNumbers := strings.Split(input, ",")

	for _, stringNumber := range stringNumbers {
		value, _ := strconv.Atoi(stringNumber)
		numbersList = append(numbersList, value)

	}
	return numbersList
}

func mapBoards(boards [][]string) (boardMap map[int]*bingoBoard, numberToMap map[int][]int) {
	boardMap = make(map[int]*bingoBoard)
	numberToMap = make(map[int][]int)
	for boardIndex, board := range boards {
		var newBoard [5][5]int
		var sum int
		for lineIndex, line := range board {

			numbersRaw := strings.Split(line, " ")
			var lineNumbers []int
			for _, stringNumber := range numbersRaw {
				if stringNumber != "" {
					value, _ := strconv.Atoi(stringNumber)
					lineNumbers = append(lineNumbers, value)
				}
			}

			for columnIndex, number := range lineNumbers {
				newBoard[lineIndex][columnIndex] = number
				numberToMap[number] = append(numberToMap[number], boardIndex)
				sum += number
			}

		}
		boardMap[boardIndex] = createBoard(newBoard, sum)

	}

	return
}
