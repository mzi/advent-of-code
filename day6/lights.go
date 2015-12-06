package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type switcher func(int) int

func read() []string {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(buf), "\n")
}

func i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parse(line string) (x, y, z, w int, verb string) {
	words := strings.Split(line, " ")
	if words[0] == "turn" {
		verb = words[1]
		words = words[1:]
	} else {
		verb = words[0]
	}
	xz := strings.Split(words[1], ",")
	yw := strings.Split(words[3], ",")
	x, z = i(xz[0]), i(xz[1])
	y, w = i(yw[0]), i(yw[1])
	return x, y, z, w, verb
}

func what(verb string, second bool) switcher {
	switch verb {
	case "on":
		if second {
			return superon
		}
		return on
	case "off":
		if second {
			return superoff
		}
		return off
	default:
		if second {
			return supertog
		}
	}
	return toggle
}

func on(i int) int  { return 1 }
func off(i int) int { return 0 }
func toggle(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
func superon(i int) int  { return i + 1 }
func supertog(i int) int { return i + 2 }
func superoff(i int) int {
	if i == 0 {
		return 0
	}
	return i - 1
}

func solve(second bool, str string, input []string) string {
	total := 0
	grid := [1000][1000]int{}
	for _, line := range input {
		x, y, z, w, verb := parse(line)
		do := what(verb, second)
		for i := x; i <= y; i++ {
			for j := z; j <= w; j++ {
				grid[i][j] = do(grid[i][j])
			}
		}
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			total += grid[i][j]
		}
	}
	return fmt.Sprintln("The total", str, "are", total)
}

func main() {
	input := read()
	solve(false, "number of lights on", input)
	solve(true, "luminacity", input)
}
