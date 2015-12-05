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
	h    *map[string]bool
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

func (s *santa) deliver(x byte) {
	h := *s.h
	s.move(x)
	if !h[s.location()] {
		h[s.location()] = true
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
	alone := map[string]bool{"0,0": true}
	shared := map[string]bool{"0,0": true}
	onlysanta := &santa{x: 0, y: 0, uniq: 1, h: &alone}
	realsanta := &santa{x: 0, y: 0, uniq: 1, h: &shared}
	robosanta := &santa{x: 0, y: 0, uniq: 0, h: &shared} // real santa gets the first house
	buf := read()
	for i, x := range buf {
		onlysanta.deliver(x)
		if i%2 == 0 {
			realsanta.deliver(x)
		} else {
			robosanta.deliver(x)
		}
	}
	fmt.Println("On his own, Santa visits", onlysanta.uniq, "houses")
	fmt.Println("With his robot friend they visit", robosanta.uniq+realsanta.uniq,
		"houses. (Santa", realsanta.uniq, "houses and robot", robosanta.uniq, "houses)")
}
