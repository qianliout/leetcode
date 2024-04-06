package main

import (
	"fmt"
)

func main() {
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 5}))
	fmt.Println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(gas) == 0 || sumPath(gas) < sumPath(cost) {
		return -1
	}
	rem := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		if rem+gas[i]-cost[i] < 0 {
			rem = 0
			start = i + 1
			continue
		}
		rem = rem + gas[i] - cost[i]
	}

	return start
}

func sumPath(num []int) int {
	ans := 0
	for _, ch := range num {
		ans += ch
	}
	return ans
}

func canCompleteCircuit2(gas []int, cost []int) int {
	total, cur, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			cur, start = 0, i+1
		}
	}
	if total < 0 {
		return -1
	}

	return start
}
