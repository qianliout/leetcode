package main

import (
	"sort"
)

func main() {

}

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := len(citations)
	for _, ch := range citations {
		if ch < h {
			h--
		}
	}
	return h
}
