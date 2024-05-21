package main

import (
	"fmt"
)

func main() {
	ss := []int{0, 1, 2, 3}
	for i := range ss {
		// fmt.Printf("pre addr %p\n", &i)
		fmt.Println("pre I1:", i)
		i++
		fmt.Println("after I1:", i)
		// fmt.Printf("after addr %p\n", &i)
	}
	fmt.Println("IIIIIIIIIII")
	for i := 0; i < len(ss); i++ {
		// fmt.Println("I2:", i)
		i++
	}
}
