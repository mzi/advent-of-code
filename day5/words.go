package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func nicep(line []byte) bool {
	v := vowels(line, 3)
	d := doubles(line)
	f := noforbidden(line)

	return v && d && f
}

func supernicep(line []byte) bool {
	p := pairs(line)
	s := skip(line)
	return p && s
}

func vowels(line []byte, o int) bool {
	x := 0
	for _, c := range line {
		switch string(c) {
		case "a", "e", "i", "o", "u":
			x++
		default:
		}
	}
	if x >= o {
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
		fmt.Println(l[i:i+2])
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
	var nice, naughty, supernice, supernaughty int
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
		if nicep(l) {
			nice++
		} else {
			naughty++
		}
		if supernicep(l) {
			supernice++
		} else {
			supernaughty++
		}
	}
	fmt.Println(nice, naughty)
	fmt.Println(supernice, supernaughty)
}
