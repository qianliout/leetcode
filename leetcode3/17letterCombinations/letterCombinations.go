package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
}

func letterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitMap := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	dd := []byte(digits)
	if len(dd) == 0 {
		return []string{}
	}
	queue := make([][]byte, 0)
	queue = append(queue, []byte{})
	for i := range dd {
		length := len(queue)
		for j := 0; j < length; j++ {
			fir := queue[0]
			queue = queue[1:]
			let := digitMap[dd[i]]

			for _, b := range let {
				th := append([]byte{}, fir...)
				th = append([]byte{}, b)
				queue = append(queue, th)
			}
		}
	}
	ans := make([]string, 0)
	for i := range queue {
		ans = append(ans, string(queue[i]))
	}

	return ans
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitMap := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	ans := make([]string, 0)
	path := make([]byte, 0)
	dfs([]byte(digits), 0, path, digitMap, &ans)
	return ans
}

func dfs(nums []byte, start int, path []byte, digitMap map[byte][]byte, res *[]string) {
	if len(path) == len(nums) {
		*res = append(*res, string(path))
		return
	}
	for _, b := range digitMap[nums[start]] {
		path = append(path, b)
		dfs(nums, start+1, path, digitMap, res)
		path = path[:len(path)-1]
	}
}
