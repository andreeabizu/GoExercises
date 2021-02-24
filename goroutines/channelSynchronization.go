package main

import "fmt"

func f1(c chan bool) {
	fmt.Println("this is before")
	c <- true
}

func f2() {
	fmt.Println("after")
}

func main() {

	c := make(chan bool)

	go f1(c)

	if <-c {
		go f2()
		fmt.Scanln()
	}

}
