package main

import (
	"fmt"
	"testing"
	"strings"
)

func scanParse (line string) (x1, x2, y1, y2 int) {
		switch {
		case strings.HasPrefix(line, "turn on"):
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			fmt.Sprint(x1, y1, x2, y2)
		case strings.HasPrefix(line, "turn off"):
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			fmt.Sprint(x1, y1, x2, y2)
		case strings.HasPrefix(line, "toggle"):
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			fmt.Sprint(x1, y1, x2, y2)
		default:
		}
	return x1, x2, y1, y2
}

func BenchmarkScanParse(b *testing.B) {
	for n := 0 ; n < b.N ; n++ {
		for _, line := range []string{"turn off 301,3 through 808,453",
			"turn on 351,678 through 951,908",
			"toggle 720,196 through 897,994"} {
			fmt.Sprint(scanParse(line))
		}
	}
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, line := range []string{"turn off 301,3 through 808,453",
			"turn on 351,678 through 951,908",
			"toggle 720,196 through 897,994"} {
			fmt.Sprint(parse(line))
		}
	}
}

func BenchmarkSolve1(b *testing.B) {
	input := read()
	b.ResetTimer()
	for n := 0 ; n < b.N ; n++ {
		_ = solve(false, "", input)
	}
}

func BenchmarkSolve2(b *testing.B) {
	input := read()
	b.ResetTimer()
	for n := 0 ; n < b.N ; n++ {
		_ = solve(true, "", input)
	}
}

/*
PASS
BenchmarkScanParse-4	   50000	     29469 ns/op
BenchmarkParse-4    	  200000	      7251 ns/op
BenchmarkSolve1-4   	      10	 131548583 ns/op
BenchmarkSolve2-4   	      10	 138284689 ns/op
ok  	github.com/mzi/advent-of-code/day6	6.288s
*/