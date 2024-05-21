package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestSubsequence("bcabc"))
}

func smallestSubsequence(s string) string {
	ss := []byte(s)
	window := make([]byte, 0)
	exit := make(map[byte]bool)
	lastIndex := make(map[byte]int)
	for i, ch := range ss {
		lastIndex[ch] = i
	}

	for i, ch := range ss {
		if exit[ch] {
			continue
		}
		for len(window) > 0 && window[len(window)-1] > ch && lastIndex[window[len(window)-1]] > i {
			exit[window[len(window)-1]] = false
			window = window[:len(window)-1]
		}
		exit[ch] = true
		window = append(window, ch)
	}
	return string(window)
}
