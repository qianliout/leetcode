package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}

func restoreIpAddresses(s string) []string {
	ss := []byte(s)
	path := make([]string, 0)
	ans := make([]string, 0)
	dfs(ss, path, 0, &ans)
	return ans

}

func dfs(nums []byte, path []string, start int, ans *[]string) {
	if len(path) > 4 {
		return
	}
	if start >= len(nums) {
		if valid(path) {
			*ans = append(*ans, strings.Join(path, "."))
		}
		return
	}

	for i := start; i < len(nums); i++ {
		path = append(path, string(nums[start:i+1]))
		dfs(nums, path, i+1, ans)
		path = path[:len(path)-1]
	}
}

func valid(path []string) bool {
	if len(path) != 4 {
		return false
	}

	for i := range path {
		if len(path[i]) > 1 && path[i][0] == '0' {
			return false
		}
		n, err := strconv.Atoi(path[i])
		if err != nil {
			return false
		}
		if n > 255 || n < 0 {
			return false
		}
	}
	return true
}
