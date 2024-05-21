package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSum(2, 3))
	fmt.Println(getSum(3, 3))
	fmt.Println(getSum(0, 3))
	fmt.Println(getSum(11, 3))
	fmt.Println(getSum(-12, 3))
}

func getSum(a int, b int) int {
	for b != 0 {
		c := a ^ b
		d := (a & b) << 1
		a, b = c, d
	}
	return a
}
