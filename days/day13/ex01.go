package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type tuple struct {
	x int
	y int
}

func createTuple(px, py int) *tuple {
	return &tuple{x: px, y: py}
}

type fold struct {
	position    int
	orientation string
}

func createFold(p int, o string) *fold {
	return &fold{position: p, orientation: o}
}

func main() {
	inputs := readInput()

	dots, folds := getDotsFolds(inputs)

	t := createTransparencyAndSum(dots)
	instructions := getFoldsInstructions(folds)

	foldX(t, instructions[0].position)

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func createMatrix(maxX, maxY int) [][]int {
	matrix := make([][]int, maxX+1)

	for i := 0; i <= maxX; i++ {
		line := make([]int, maxY+1, maxY+1)
		matrix[i] = line
	}

	return matrix

}

func foldY(transparency [][]int, foldPosition int) [][]int {
	foldedTransparency := make([][]int, 0)
	sum := 0
	for j, n := 0, len(transparency[0])-1; j < foldPosition; j, n = j+1, n-1 {
		line := make([]int, 0)
		for i := range transparency {
			if transparency[i][j]+transparency[i][n] > 0 {
				line = append(line, 1)
				sum++
			} else {
				line = append(line, 0)
			}
		}
		foldedTransparency = append(foldedTransparency, line)
	}

	fmt.Println(sum)

	return foldedTransparency

}

func foldX(transparency [][]int, foldPosition int) [][]int {
	foldedTransparency := make([][]int, 0)
	sum := 0
	for i, m := 0, len(transparency)-1; i < foldPosition; i, m = i+1, m-1 {
		line := make([]int, 0)
		for j := range transparency[i] {
			if transparency[i][j]+transparency[m][j] > 0 {
				line = append(line, 1)
				sum++
			} else {
				line = append(line, 0)
			}
		}
		foldedTransparency = append(foldedTransparency, line)
	}

	fmt.Println(sum)

	return foldedTransparency

}

func createTransparencyAndSum(dots []string) [][]int {
	tuples := make([]*tuple, 0)
	sum := 0
	maxY := 0
	maxX := 0

	for _, dotLine := range dots {
		stringDots := strings.Split(dotLine, ",")
		if len(stringDots) > 1 {
			x, _ := strconv.Atoi(string(stringDots[0]))
			y, _ := strconv.Atoi(string(stringDots[1]))

			maxX = getMax(maxX, x)
			maxY = getMax(maxY, y)

			dotTuple := createTuple(x, y)
			tuples = append(tuples, dotTuple)
		}

	}

	transparency := createMatrix(maxX, maxY)

	for _, dotTuple := range tuples {
		transparency[dotTuple.x][dotTuple.y] = 1
		sum++
	}

	return transparency

}

func getFoldsInstructions(foldsString []string) []*fold {
	foldsInstructions := make([]*fold, 0)
	for _, v := range foldsString {
		v = v[11:]
		foldData := strings.Split(v, "=")
		foldPosition, _ := strconv.Atoi(string(foldData[1]))
		foldOrientation := foldData[0]
		newFold := createFold(foldPosition, foldOrientation)

		foldsInstructions = append(foldsInstructions, newFold)
	}

	return foldsInstructions
}

func getDotsFolds(inputs []string) (dots, folds []string) {

	onFolds := false

	for _, value := range inputs {
		if len(value) == 0 {
			onFolds = true
		} else {
			if onFolds {
				folds = append(folds, value)
			} else {
				dots = append(dots, value)
			}
		}
	}
	return
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day13/input.txt")
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

func convertStringLineToNumbers(input string) []int {
	var numbersList []int

	for _, stringNumber := range input {
		value, _ := strconv.Atoi(string(stringNumber))
		numbersList = append(numbersList, value)

	}
	return numbersList
}
