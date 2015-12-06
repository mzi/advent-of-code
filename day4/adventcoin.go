package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

const secretCode = "ckczppom"

func sum(ch5, ch6 chan<- uint64, secret <-chan uint64) {
	for x := range secret {
		s := fmt.Sprintf("%s%d", secretCode, x)
		md5sum := fmt.Sprintf("%x", md5.Sum([]byte(s)))
		if strings.HasPrefix(md5sum, "00000") {
			ch5 <- x
		}
		if strings.HasPrefix(md5sum, "000000") {
			ch6 <- x
			break
		}
	}
}

func incrementer(secret chan<- uint64) {
	var x uint64
	for x = 0; x < math.MaxUint64; x++ {
		secret <- x
	}
}

func main() {
	ch5 := make(chan uint64)
	ch6 := make(chan uint64)
	secret := make(chan uint64)

	for i := 0; i <= 12; i++ {
		go sum(ch5, ch6, secret)
	}

	go incrementer(secret)

	for {
		select {
		case secret5 := <-ch5:
			fmt.Println("Five zeroes has secret code", secret5)
		case secret6 := <-ch6:
			fmt.Println("Six zeroes has secret code", secret6)
		}
	}
}
