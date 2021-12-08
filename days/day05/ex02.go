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

type arrow struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func createArrow(x1, y1, x2, y2 int) *arrow {
	a := arrow{x1: x1, y1: y1, x2: x2, y2: y2}
	return &a
}

func main() {
	inputs := readInput()
	arrows, maxX, maxY := createArrowsAndTableSize(inputs)

	board := createBoard(maxX, maxY)

	fmt.Println("--------------------")
	overlaps := getOverlaps(arrows, board)
	fmt.Println(overlaps)

}

func getMinMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func getOverlaps(arrows []*arrow, board [][]int) int {
	var overlays int

	for _, arrow := range arrows {
		if arrow.x1 == arrow.y2 && arrow.y1 == arrow.x2 {
			min, max := getMinMax(arrow.x1, arrow.x2)

			for i, j := min, max; i <= max; i, j = i+1, j-1 {
				board[i][j] = board[i][j] + 1

				if board[i][j] == 2 {
					overlays++
				}
			}
		} else if arrow.x1 == arrow.x2 && arrow.y1 != arrow.y2 {
			xPosition := arrow.x1
			minY, maxY := getMinMax(arrow.y1, arrow.y2)
			for j := minY; j <= maxY; j++ {
				board[j][xPosition] = board[j][xPosition] + 1

				if board[j][xPosition] == 2 {
					overlays++
				}
			}
		} else if arrow.y1 == arrow.y2 && arrow.x1 != arrow.x2 {
			yPosition := arrow.y1
			minX, maxX := getMinMax(arrow.x1, arrow.x2)

			for i := minX; i <= maxX; i++ {

				board[yPosition][i] = board[yPosition][i] + 1

				if board[yPosition][i] == 2 {
					overlays++
				}
			}
		} else if arrow.x1 == arrow.y1 && arrow.x2 == arrow.y2 {
			min, max := getMinMax(arrow.x1, arrow.x2)

			for i := min; i <= max; i++ {
				board[i][i] = board[i][i] + 1

				if board[i][i] == 2 {
					overlays++
				}
			}

		} else if arrow.x1-arrow.x2 == arrow.y1-arrow.y2 {
			deltaX := arrow.x1 - arrow.x2

			if deltaX < 0 {
				arrow.x1, arrow.x2, arrow.y1, arrow.y2 = arrow.x2, arrow.x1, arrow.y2, arrow.y1
				deltaX = -deltaX
			}

			for i, j := arrow.x2, arrow.y2; i <= arrow.x1; i, j = i+1, j+1 {
				board[j][i] = board[j][i] + 1
				if board[j][i] == 2 {
					overlays++
				}
			}

		} else if arrow.x1-arrow.x2 == -(arrow.y1 - arrow.y2) {
			deltaX := arrow.x1 - arrow.x2

			if deltaX > 0 {
				arrow.x1, arrow.x2, arrow.y1, arrow.y2 = arrow.x2, arrow.x1, arrow.y2, arrow.y1
				deltaX = -deltaX
			}

			for i, j := arrow.x1, arrow.y1; i <= arrow.x2; i, j = i+1, j-1 {
				board[j][i] = board[j][i] + 1
				if board[j][i] == 2 {
					overlays++
				}
			}

		}
	}

	return overlays

}

func createBoard(maxX, maxY int) [][]int {
	matrix := make([][]int, maxX+1)

	for i := 0; i <= maxX; i++ {
		line := make([]int, maxY+1, maxY+1)
		matrix[i] = line
	}

	return matrix

}

func extractPoints(pointString string) (x, y int) {
	pointString = strings.TrimSpace(pointString)
	coordinatesXY := strings.Split(pointString, ",")

	x, _ = strconv.Atoi(coordinatesXY[0])
	y, _ = strconv.Atoi(coordinatesXY[1])
	return
}

func createArrowsAndTableSize(input []string) (arrowList []*arrow, maxX, maxY int) {
	for _, line := range input {
		coordinates := strings.Split(line, "->")
		beginString := coordinates[0]
		endString := coordinates[1]

		x1, y1 := extractPoints(beginString)
		x2, y2 := extractPoints(endString)

		maxX, maxY = updateMax(x1, x2, maxX, y1, y2, maxY)

		arrow := createArrow(x1, y1, x2, y2)
		arrowList = append(arrowList, arrow)
	}
	return
}

func updateMax(x1, x2, xf, y1, y2, yf int) (int, int) {
	xLine := []int{x1, x2, xf}
	yLine := []int{y1, y2, yf}

	sort.Ints(xLine)
	sort.Ints(yLine)

	return xLine[2], yLine[2]

}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day05/input.txt")
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
