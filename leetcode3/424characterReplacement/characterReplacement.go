package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement2(s string, k int) int {
	window := make(map[uint8]int)

	ans := 0
	le, ri := 0, 0
	// 双指针
	for le <= ri && ri < len(s) {
		window[s[ri]]++
		ri++
		// 一定要先入窗口，再检测数据，再更新结果，
		// 如果先检查结果，再入窗口的话，最后一个结果会更新不到
		for !check(window, k) {
			window[s[le]]--
			le++
		}
		sm := sum(window)
		if ans < sm {
			ans = sm
		}
	}

	return ans
}

func check(data map[uint8]int, k int) bool {
	mx := 0
	sm := 0
	for _, ch := range data {
		sm += ch
		if ch > mx {
			mx = ch
		}
	}
	return mx+k >= sm
}

func sum(data map[uint8]int) int {
	sm := 0
	for _, ch := range data {
		sm += ch
	}
	return sm
}

func characterReplacement(s string, k int) int {
	exit := make(map[uint8]int)
	le, ri, ma, ans := 0, 0, 0, 0
	for le <= ri && ri < len(s) {
		exit[s[ri]]++
		if exit[s[ri]] > ma {
			ma = exit[s[ri]]
		}
		ri++

		for le <= ri && ri-le > ma+k {
			// 这里更新了 exit 之后，为啥不更新 ma 的值，是个难点
			exit[s[le]]--
			le++
		}
		if ri-le > ans {
			ans = ri - le
		}
	}
	return ans
}
