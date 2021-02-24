package main

import "fmt"

func sendValues(c chan string) {
	c <- "ana"
	c <- "maria"
	c <- "ioana"
	c <- "paul"
	c <- "matei"
	close(c)
}

func main() {
	c := make(chan string, 5)

	go sendValues(c)

	for el := range c {
		fmt.Println(el)
	}
}
