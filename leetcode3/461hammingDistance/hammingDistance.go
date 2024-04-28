package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(3, 1))
}

func hammingDistance(x int, y int) int {
	cnt := 0
	for x != y {
		if x&1 != y&1 {
			cnt += 1
		}
		x = x >> 1
		y = y >> 1
	}

	return cnt
}
