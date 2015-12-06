package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func niceWords(line []byte) bool {
	v := vowels(line)
	d := doubles(line)
	f := noforbidden(line)
	return v && d && f
}

func nicerWords(line []byte) bool {
	p := pairs(line)
	s := skip(line)
	return p && s
}

func vowels(line []byte) bool {
	n := 0
	for _, char := range line {
		switch string(char) {
		case "a", "e", "i", "o", "u":
			n++
		default:
		}
	}
	if n >= 3 {
		return true
	}
	return false
}

func doubles(line []byte) bool {
	for i := 1; i < len(line); i++ {
		if line[i-1] == line[i] {
			return true
		}
	}
	return false
}

func noforbidden(line []byte) bool {
	l := string(line)
	for _, c := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(l, c) {
			return false
		}
	}
	return true
}

func pairs(line []byte) bool {
	l := string(line)
	for i := 0 ; i < len(l) - 2 ; i++ {
		if strings.Contains(l[i+2:], l[i:i+2]) {
			return true
		}
	}
	return false
}

func skip(line []byte) bool {
	for i := len(line) - 1; i > 2; i-- {
		if line[i] == line[i-2] {
			return true
		}
	}
	return false
}

func main() {
	var nice, nicer int
	fd, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := string(scanner.Text())
		l := []byte(line)
		if niceWords(l) {
			nice++
		}
		if nicerWords(l) {
			nicer++
		}
	}
	fmt.Println("There are", nice, "words.")
	fmt.Println("There are", nicer, "even nicer words.")
}
