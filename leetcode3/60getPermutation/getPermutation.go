package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getPermutation(3, 4))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
}

func getPermutation(n int, k int) string {

	used := make(map[int]bool)
	mem := make(map[int]int)
	if k > fib(n, mem) {
		return ""
	}

	ans := make([]string, 0)
	for k > 1 {
		cnt := help(n, used, &k, mem)
		ans = append(ans, strconv.Itoa(cnt))
	}
	for i := 1; i <= n; i++ {
		if !used[i] {
			ans = append(ans, strconv.Itoa(i))
		}
	}

	return strings.Join(ans, "")
}

func help(n int, used map[int]bool, k *int, mem map[int]int) int {
	i := 1
	for i = 1; i <= n; i++ {
		if used[i] {
			continue
		}
		cnt := fib(n-1-len(used), mem)
		if *k > cnt {
			*k = *k - cnt
		} else {
			// 这里要做一步处理
			break
		}
	}
	used[i] = true
	return i
}

func fib(n int, exit map[int]int) int {
	if n <= 1 {
		return 1
	}
	if va, ok := exit[n]; ok {
		return va
	}
	va := n * fib(n-1, exit)
	exit[n] = va
	return va
}
