package main

/*

Second day problem.

Find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l.
The elves also need a little extra paper for each present: the
area of the smallest side.

The ribbon required to wrap a present is the shortest distance
around its sides, or the smallest perimeter of any one face. Each
present also requires a bow made out of ribbon as well; the feet
of ribbon required for the perfect bow is equal to the cubic feet
of volume of the present.

*/


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func paper(l, w, h int) int {
	lw, hl, wh := l*w, h*l, w*h
	s := []int{lw, hl, wh}
	sort.Ints(s)
	return 2*lw + 2*hl + 2*wh + s[0]
}

func ribbon(l, w, h int) int {
	s := []int{l, w, h}
	sort.Ints(s)
	return 2*s[0] + 2*s[1] + l*w*h
}

func main() {
	sqFeet := 0
	feet := 0

	fd, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := string(scanner.Text())
		geo := strings.Split(line, "x")
		l, _ := strconv.Atoi(geo[0])
		w, _ := strconv.Atoi(geo[1])
		h, _ := strconv.Atoi(geo[2])

		sqFeet += paper(l, w, h)
		feet += ribbon(l, w, h)
	}

	fmt.Println("Paper:", sqFeet, "sqfeet", "Ribbon:", feet, "ft")
}
