package main

func main() {

}

func canMakePaliQueries1(s string, queries [][]int) []bool {
	preSum := make([][26]int, len(s)+1) // 前缀和，统计前面有多少个奇数
	for i, c := range s {
		preSum[i+1] = preSum[i]
		preSum[i+1][c-'a']++
	}
	ans := make([]bool, len(queries))
	for i, qa := range queries {
		le, ri, k, m := qa[0], qa[1], qa[2], 0
		for j := 0; j < 26; j++ {
			// 这样写是不对的 比如：(5-3)%2==1，正确的答案应该是：5%2-3%2 = 0
			// m = m + (preSum[ri+1][j]%2 - preSum[le][j]%2)
			m = m + (preSum[ri+1][j]-preSum[le][j])%2
		}
		ans[i] = m/2 <= k
	}
	return ans
}

// 只关心奇偶性
func canMakePaliQueries2(s string, queries [][]int) []bool {
	preSum := make([][26]int, len(s)+1) // 前缀和，统计每个字母前面有多少个奇数
	for i, c := range s {
		preSum[i+1] = preSum[i]
		preSum[i+1][c-'a']++
		preSum[i+1][c-'a'] %= 2
	}
	ans := make([]bool, len(queries))
	for i, qa := range queries {
		le, ri, k, m := qa[0], qa[1], qa[2], 0
		for j := 0; j < 26; j++ {
			if preSum[ri+1][j] != preSum[le][j] {
				m++
			}
		}
		ans[i] = m/2 <= k
	}
	return ans
}

// 使用异或运算
func canMakePaliQueries(s string, queries [][]int) []bool {
	preSum := make([][26]int, len(s)+1) // 前缀和，统计前面有多少个奇数
	for i, c := range s {
		preSum[i+1] = preSum[i]
		// 如果preSum[i+1][c-'a']原来是0，现在变成1，如果原来是1，现在变成0
		preSum[i+1][c-'a'] ^= 1
	}
	ans := make([]bool, len(queries))
	for i, qa := range queries {
		le, ri, k, m := qa[0], qa[1], qa[2], 0
		for j := 0; j < 26; j++ {
			if preSum[ri+1][j]^preSum[le][j] == 1 {
				m++
			}
		}
		ans[i] = m/2 <= k
	}
	return ans
}
