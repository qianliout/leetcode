package main

func main() {

}

/*
编写一段程序来查找第 n 个超级丑数。
超级丑数是指其所有质因数都是长度为 k 的质数列表 primes 中的正整数。
示例:
输入: n = 12, primes = [2,7,13,19]
输出: 32
解释: 给定长度为 4 的质数列表 primes = [2,7,13,19]，前 12 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
说明:

	1 是任何给定 primes 的超级丑数。
	 给定 primes 中的数字以升序排列。
	0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000 。
	第 n 个超级丑数确保在 32 位有符整数范围内。
*/
func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	if len(primes) == 0 {
		return 1
	}

	return 0
}
