package main

import "os"

func main() {
	file, _ := os.Create("file.txt")
	defer file.Close()

	file.WriteString("andreea")

}
