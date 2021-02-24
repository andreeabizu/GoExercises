package main

import "fmt"

func f1(c1 chan int, c2 chan string) {
	x := 10
	for {
		select {
		case c1 <- x:
			x--
		case <-c2:
			fmt.Println("done")
			return
		}
	}
}

func f2(c1 chan int, c2 chan string) {

	for i := 0; i <= 5; i++ {
		fmt.Println(<-c1)
	}
	c2 <- "done"
}

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan string)
	go f1(c1, c2)
	go f2(c1, c2)
	fmt.Scanln()
}
