package main

// Much shorter but slower solution to the first elevator problem.

import (
	"fmt"
	"regexp"
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
	buf    := read()
	up   := regexp.MustCompile("\\(").FindAll(buf, -1)
	down := regexp.MustCompile("\\)").FindAll(buf, -1)
	fmt.Println("Floor is", len(up) - len(down))
}
