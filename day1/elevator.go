package main

/*

First day problem

An opening parenthesis, (, means he should go up one floor, and a closing
parenthesis, ), means he should go down one floor.

To what floor do the instructions take Santa?

Now, given the same instructions, find the position of the first character
that causes him to enter the basement (floor -1). The first character in
the instructions has position 1, the second character has position 2, and
so on.

*/

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
	first := true
	for i, direction := range input {
		switch string(direction) {
		case "(":
			floor++
		case ")":
			floor--
		default:
		}
		if floor == -1 && first {
			fmt.Println("Santa visits the basement for the first time at pos ",i + 1)
			first = false
		}
	}
	fmt.Println("Santa ends up on floor", floor)
}
