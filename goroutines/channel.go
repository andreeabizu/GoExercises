package main

import "fmt"

func f1(s chan string) {
	s <- "message from f1"
}

func f2(s chan string) {

	x := <-s
	fmt.Println(x)
}

func main() {
	var c chan string = make(chan string)
	go f1(c)
	go f2(c)

	fmt.Scanln()
}
