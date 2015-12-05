package main

/*

Moves are always exactly one house to the north (^), south (v), east (>),
or west (<). After each move, he delivers another present to the house at
his new location.

How many houses receive at least one present?

Santa and Robo-Santa start at the same location (delivering two presents to
the same starting house), then take turns moving based on the same instructions
as before

How many houses receive at least one present?

*/

import (
	"fmt"
	"io/ioutil"
	"log"
)

type santa struct {
	x    int
	y    int
	uniq int
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
	}
}

func (s *santa) location() string {
	return fmt.Sprintf("%d,%d", s.x, s.y)
}

func (s *santa) deliver(x byte, h *map[string]bool) {
	house := *h
	s.move(x)
	if !house[s.location()] {
		house[s.location()] = true
		s.uniq++
	}
}

func read() []byte {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func main() {
	onlysanta := new(santa)
	realsanta := new(santa)
	robosanta := new(santa)
	onlysanta.x, onlysanta.y, onlysanta.uniq = 0, 0, 1
	realsanta.x, realsanta.y, realsanta.uniq = 0, 0, 1
	robosanta.x, robosanta.y = 0, 0
	buf := read()
	alone := map[string]bool{"0,0": true}
	shared := map[string]bool{"0,0": true}
	for i, x := range buf {
		onlysanta.deliver(x, &alone)
		if i%2 == 0 {
			realsanta.deliver(x, &shared)
		} else {
			robosanta.deliver(x, &shared)
		}
	}
	fmt.Println("On his own, Santa visits", onlysanta.uniq, "houses")
	fmt.Println("With his robot friend they visit", robosanta.uniq+realsanta.uniq,
		"houses. (Santa", realsanta.uniq, "houses and robot", robosanta.uniq, "houses)")
}
