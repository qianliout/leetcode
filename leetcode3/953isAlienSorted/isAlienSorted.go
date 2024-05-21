package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
}

func isAlienSorted(words []string, order string) bool {
	ord := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		ord[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		if !compare(words[i], words[i+1], ord) {
			return false
		}
	}
	return true
}

func compare(a, b string, ord map[byte]int) bool {
	le := len(a)
	if len(b) < len(a) {
		le = len(b)
	}
	for j := 0; j < le; j++ {
		if ord[a[j]] < ord[b[j]] {
			return true
		} else if ord[a[j]] > ord[b[j]] {
			return false
		}
	}
	return len(a) <= len(b)
}
