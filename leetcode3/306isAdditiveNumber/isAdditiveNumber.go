package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("112358"))
}

func isAdditiveNumber(num string) bool {
	var res bool

	backTrack([]byte(num), 0, 0, 0, 0, &res)
	return res
}

func backTrack(num []byte, cnt, start, fir, sec int, res *bool) {
	if *res {
		return
	}
	for i := start; i < len(num); i++ {
		// 不可以有01,这种情况
		if num[start] == '0' && i > start {
			continue
		}
		n, _ := strconv.Atoi(string(num[start : i+1]))

		if cnt >= 2 {
			if n == fir+sec {
				backTrack(num, cnt+1, i+1, sec, n, res)
			}
		} else {
			backTrack(num, cnt+1, i+1, sec, n, res)
		}
	}
}
