package main

import "fmt"

func f1(c chan int) {

	for i := 0; i <= 10; i++ {
		c <- i
	}
}

func f2(c chan string) {
	c <- "done"
}

func f3(c1 chan int, c2 chan string) {

	for {
		select {
		case x := <-c1:
			fmt.Println(x)
		case y := <-c2:
			fmt.Println(y)

		}
	}
}

func main() {

	c1 := make(chan int, 11)
	c2 := make(chan string)

	go f3(c1, c2)
	go f1(c1)
	go f2(c2)

	fmt.Scanln()

}
