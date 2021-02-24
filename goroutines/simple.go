package main

import "fmt"

func print(s string) {

	fmt.Println(s)
}

func main() {
	print("before goroutine")
	go print("in goroutine")

	fmt.Scanln()
}
