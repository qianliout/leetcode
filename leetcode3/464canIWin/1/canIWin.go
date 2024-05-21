package main

import (
	"fmt"
)

func main() {
	fmt.Println(canIWin(10, 11)) // false
	fmt.Println(canIWin(10, 0))  // true
	fmt.Println(canIWin(10, 1))  // true
	fmt.Println(canIWin(18, 79)) // true
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	choose := make(map[int]bool)
	for i := 1; i <= maxChoosableInteger; i++ {
		choose[i] = false
	}
	return dfs(choose, 0, desiredTotal)
}

// timeout
// 可以得到结果，但是会超时
func dfs(choose map[int]bool, sum int, desired int) bool {
	for k := range choose {

		if sum+k >= desired {
			return true
		}
		co := Copy(choose, k)
		if !dfs(co, sum+k, desired) {
			return true
		}
	}
	return false
}

func Copy(choose map[int]bool, val int) map[int]bool {
	ans := make(map[int]bool)
	for k, v := range choose {
		if k != val {
			ans[k] = v
		}
	}
	return ans
}
