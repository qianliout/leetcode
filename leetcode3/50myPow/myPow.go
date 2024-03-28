package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if x == 1 {
		return x
	}
	half := myPow(x, n>>1)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}
