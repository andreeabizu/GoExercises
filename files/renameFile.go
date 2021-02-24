package main

import (
	"os"
)

func rename(src, dst string) {
	os.Rename(src, dst)
}

func main() {
	rename("file.txt", "andreea.txt")
}
