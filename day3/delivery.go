package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type santa struct {
	x int
	y int
}

func (s *santa) move(b byte) {
	switch string(b) {
	case "^":
		s.x++
	case "v":
		s.x--
	case ">":
		s.y++
	case "<":
		s.y--
	default:
		fmt.Println("Uknown byte!", string(b))
	}
}

func (s *santa) location() string {
	return fmt.Sprintf("%d,%d", s.x, s.y)
}

func read() []byte {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func main() {
	santa := new(santa)
	santa.x, santa.y = 0, 0
	buf := read()
	delivery := 1
	houses := map[string]bool{"0,0": true}
	for _, x := range buf {
		santa.move(x)
		if !houses[santa.location()] {
			houses[santa.location()] = true
			delivery++
		}
	}
	fmt.Println(delivery)
}
