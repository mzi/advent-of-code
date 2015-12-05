package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calc(l, w, h int) int {
	lw := l * w
	hl := h * l
	wh := w * h
	s := []int{lw, hl, wh}
	sort.Ints(s)
	return 2*lw + 2*hl + 2*wh + s[0]
}

func main() {
	sqFeet := 0

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

		sqFeet += calc(l, w, h)
	}

	fmt.Println(sqFeet)
}
