package main

import "fmt"

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	x := factorial(4)
	y := factorial(5)
	z := factorial(3)

	fmt.Println(x, y, z)
}
