package main

import "fmt"

func helloWorld() (string, string) {
	return "hello", "world"
}

func main() {
	x, y := helloWorld()

	fmt.Println(x)
	fmt.Println(y)
}
