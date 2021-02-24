package main

import "fmt"

func f1(c chan<- string) {
	c <- "first"
	c <- "second"
	c <- "last"
	c <- "gg"
}

func f2(c <-chan string) {
	fmt.Println(<-c, <-c, <-c)
}

func main() {
	var c chan string = make(chan string, 3)
	go f1(c)
	go f2(c)
	fmt.Println("in main", <-c)
	fmt.Scanln()
}
