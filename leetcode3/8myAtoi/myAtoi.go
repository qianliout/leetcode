package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("43"))
	fmt.Println(myAtoi("-43"))
	fmt.Println(myAtoi("-43a"))
	fmt.Println(myAtoi("-430a"))
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	ss := []byte(s)
	num, pos := 0, 1
	for i, b := range ss {
		if i == 0 && b == '+' {
			continue
		}
		if i == 0 && b == '-' {
			pos = -1
			continue
		}
		if !dig(b) {
			break
		}

		n := int(b - '0')

		if num > (math.MaxInt32-n)/10 && pos == 1 {
			return math.MaxInt32
		}

		if -num < (math.MinInt32+n)/10 && pos == -1 {
			return math.MinInt32
		}
		num = num*10 + n
	}
	return num * pos
}

func dig(s byte) bool {
	return s >= '0' && s <= '9'
}
