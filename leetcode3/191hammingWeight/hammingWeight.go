package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(128))

}

func hammingWeight(n int) int {
	ans := 0
	for n > 0 {
		if n&1 == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}
