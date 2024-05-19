package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPalindrome(2))
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	mx := int(math.Pow(10, float64(n))) - 1
	for i := mx; i >= 0; i-- {
		num := i
		t := i
		for t != 0 {
			num = num*10 + (t % 10)
			t = t / 10
		}
		for j := mx; j*j >= num; j-- {
			if num%j == 0 {
				return num % 1337
			}
		}
	}
	return -1
}
