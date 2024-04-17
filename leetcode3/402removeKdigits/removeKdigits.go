package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("9", 1))
	fmt.Println(removeKdigits("112", 1))
}

func removeKdigits(num string, k int) string {
	window := make([]byte, 0)
	remin := len(num) - k
	for _, ch := range num {
		for k > 0 && len(window) > 0 && window[len(window)-1] > byte(ch) {
			k--
			window = window[:len(window)-1]
		}
		window = append(window, byte(ch))
	}
	// 而当遍历完成，如果 k 仍然大于 0。不妨假设最终还剩下 x 个需要丢弃，那么我们需要选择删除末尾 x 个元素。
	// 题目要求我们丢弃 k 个。反过来说，不就是让我们保留n-k个元素么？其中 n 为数字长度。
	// 那么我们只需要按照上面的方法遍历完成之后，再截取前 n - k 个元素即可
	window = window[:remin]
	for len(window) > 0 {
		if window[0] == '0' {
			window = window[1:]
		} else {
			break
		}
	}
	if len(window) == 0 {
		return "0"
	}
	return string(window)
}
