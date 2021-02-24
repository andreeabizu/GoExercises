package main

import "fmt"

func sum(x, y int) (z int) {

	z = x + y
	return
}

func main() {

	var x, y = 4, 5
	fmt.Println(sum(x, y))

}
