package main

import "fmt"

func even() func() int {
	x := 0
	return func() int {
		x = x + 2
		return x
	}
}

func a(fun func(string)) {
	fun("in function a")
}

func main() {

	aFunc := func(s string) {
		fmt.Println(s)
	}

	aFunc("Something")

	a(aFunc)
	b := even()

	for i := 0; i <= 6; i++ {
		fmt.Println(b())
	}
}
