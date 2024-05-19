package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(smallestGoodBase("13"))
}

func smallestGoodBase(n string) string {
	mx, _ := strconv.Atoi(n)
	le, ri := 1, mx+1
	for le < ri {
		mid := le + (ri-le)>>1
		// 相当于左边界
		if check(mx, mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return strconv.Itoa(le)
}

// n 是值，ba 是值 位数

func check(n, ba int) bool {
	for n > 0 {
		if n%ba != 0 {
			return false
		}
		n = n / ba
	}

	return true
}
