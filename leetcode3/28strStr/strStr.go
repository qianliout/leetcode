package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("abc", "c"))
}

func strStr(haystack string, needle string) int {
	ha := []byte(haystack)
	for i := 0; i < len(haystack); i++ {
		if len(needle)+i > len(haystack) {
			break
		}
		ss := ha[i : len(needle)+i]
		if string(ss) == needle {
			return i
		}
	}
	return -1
}
