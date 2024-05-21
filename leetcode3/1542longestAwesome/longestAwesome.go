package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestAwesome("3242415"))
}

func longestAwesome(s string) int {
	exit := make(map[byte]int)
	for _, ch := range s {
		exit[byte(ch)]++
	}

	le, ri := 0, 0
	ans := 0
	for i := 1; i < len(s); i++ {
		for 

	}

	return ans
}

func check(exit map[byte]int) bool {
	hasOdd := false
	for _, ch := range exit {
		if ch%2 == 0 {
			continue
		} else if ch%2 == 1 && !hasOdd {
			hasOdd = true
			continue
		} else if ch%2 == 1 && hasOdd {
			return false
		}
	}
	return true
}
