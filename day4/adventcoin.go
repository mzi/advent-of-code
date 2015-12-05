package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const secret = "ckczppom"

func sum(x int64) string {
	s := fmt.Sprintf("%s%d", secret, x)
	m := md5.Sum([]byte(s))
	p := fmt.Sprintf("%x", m)
	return p
}

func main() {
	var x int64 = 0
	fz := false
	sz := false
	for {
		s := sum(x)
		switch {
		case strings.HasPrefix(s, "00000"):
			if !fz {
				fmt.Println(x)
				fz = true
			}
		case strings.HasPrefix(s, "000000"):
			if !sz {
				fmt.Println(x)
				sz = true
			}
		default:
		}
		if fz && sz {
			break
		}
		x++
	}
}
