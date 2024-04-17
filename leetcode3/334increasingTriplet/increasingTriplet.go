package main

import (
	"math"
)

func main() {

}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	fir, sec := nums[0], math.MaxInt64
	for _, ch := range nums {
		if ch < fir {
			fir = ch
		} else if ch > fir && ch < sec {
			sec = ch
		} else if ch > sec {
			return true
		}
	}

	return false
}
