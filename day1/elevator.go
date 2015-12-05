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
	first := true
	for i, d := range input {
		switch string(d) {
		case up:
			floor++
		case down:
			floor--
		default:
		}
		if floor == -1 && first {
			fmt.Println(i + 1)
			first = false
		}
	}
	fmt.Println(floor)
}
