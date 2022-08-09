package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type node struct {
	value     string
	neighbors []string
	nodeType  string
}

func main() {
	inputs := readInput()

	// for _, in := range inputs {
	// 	fmt.Println(in)
	// }
	nodeMap := createNodes(inputs)
	x := walk(nodeMap, "start", "")

	for i, v := range x {
		fmt.Println(v, i)
	}

}

func walk(nodeMap map[string]*node, currentNodeValue, path string) string {
	currentNode := nodeMap[currentNodeValue]
	value := currentNode.value
	if value == "end" {
		fmt.Println(path + "end")
		return path + "end"
	} else {
		newPath := path + currentNodeValue + ","
		for _, neighborValue := range currentNode.neighbors {
			if neighborValue != "start" {
				if !strings.Contains(newPath, neighborValue) || nodeMap[neighborValue].nodeType == "big" {
					walk(nodeMap, neighborValue, newPath)
				}
			}
		}

	}
	return ""
}

func checkIfIsLowerCave(nodeName string) bool {
	firstRune := rune(nodeName[0])
	return unicode.IsLower(firstRune)
}

func createNewNode(nodeName string) *node {
	nodeType := "big"
	if nodeName == "start" {
		nodeType = "start"
	} else if nodeName == "end" {
		nodeType = "end"
	} else if checkIfIsLowerCave(nodeName) {
		nodeType = "small"
	}
	n := node{value: nodeName, neighbors: make([]string, 0), nodeType: nodeType}

	return &n
}

func createNodes(inputs []string) map[string]*node {
	nodeMap := make(map[string]*node)
	for _, line := range inputs {
		places := strings.Split(line, "-")
		for _, place := range places {
			if nodeMap[place] == nil {
				newNode := createNewNode(place)
				nodeMap[place] = newNode
			}
		}

		nodeMap[places[0]].neighbors = append(nodeMap[places[0]].neighbors, places[1])
		nodeMap[places[1]].neighbors = append(nodeMap[places[1]].neighbors, places[0])
	}
	// for key, value := range nodeMap {
	// 	fmt.Println(key, value)
	// }

	return nodeMap
}

func readInput() []string {
	var input []string
	file, err := os.Open("./inputs/day12/input.txt")
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
