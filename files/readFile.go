package main

import (
	"fmt"
	"io/ioutil"
)

func read(filename string) (string, error) {

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	s := string(file)

	return s, nil
}

func main() {

	s, err := read("file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)

}
