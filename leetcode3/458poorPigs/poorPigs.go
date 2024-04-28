package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(poorPigs(1000, 15, 60))
	fmt.Println(poorPigs(4, 15, 15))
	fmt.Println(poorPigs(4, 15, 30))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets <= 0 || minutesToDie <= 0 || minutesToTest <= 0 {
		return 0
	}
	bit := minutesToTest/minutesToDie + 1
	cnt := 0
	m := 0

	for {
		m = int(math.Pow(float64(bit), float64(cnt)))
		if m >= buckets {
			return cnt
		}
		cnt++
	}
}
