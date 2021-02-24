package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLine(filepath string) ([]string, error) {

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	var lines []string
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLine("file.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}
