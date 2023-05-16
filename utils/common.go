package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadInput(filePath string) []string {
	var input []string
	file, err := os.Open(filePath)
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
