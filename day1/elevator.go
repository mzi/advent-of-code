package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func read() []byte {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func main() {
	floor := 0
	input := read()
	up := "("
	down := ")"
	for _, d := range input {
		switch string(d) {
		case up:
			floor++
		case down:
			floor--
		default:
			fmt.Println("We shouldn't be here")
		}
	}
	fmt.Println(floor)
}
