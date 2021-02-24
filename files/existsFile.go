package main

import (
	"fmt"
	"os"
)

func main() {

	file := "file.txt"
	if _, err := os.Stat(file); err == nil {
		fmt.Printf("%s exists", file)
	} else {
		fmt.Printf("%s don't exist", file)
	}
}
