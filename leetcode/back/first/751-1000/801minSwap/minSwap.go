package main

import (
	"outback/leetcode/back/common"
)

func main() {

}

/*
我们有两个长度相等且不为空的整型数组 A 和 B 。
我们可以交换 A[i] 和 B[i] 的元素。注意这两个元素在各自的序列中应该处于相同的位置。
在交换过一些元素之后，数组 A 和 B 都应该是严格递增的（数组严格递增的条件仅为A[0] < A[1] < A[2] < ... < A[A.length - 1]）。
给定数组 A 和 B ，请返回使得两个数组均保持严格递增状态的最小交换次数。假设给定的输入总是有效的。
示例:
输入: A = [1,3,5,4], B = [1,2,3,7]
输出: 1
解释:
交换 A[3] 和 B[3] 后，两个数组如下:
A = [1, 3, 5, 7] ， B = [1, 2, 3, 4]
两个数组均为严格递增的。
注意:
A, B 两个数组的长度总是相等的，且长度的范围为 [1, 1000]。
A[i], B[i] 均为 [0, 2000]区间内的整数。
*/
func minSwap(A []int, B []int) int {
	dp1 := make(map[int]int) // 不交换
	dp2 := make(map[int]int) // 交换

	dp1[0] = 0
	dp2[0] = 1
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] && B[i-1] < B[i] {
			// 任意交换或者不交换，取最优值
			if A[i-1] < B[i] && B[i-1] < A[i] {
				dp1[i] = common.Min(dp1[i-1], dp2[i-1])     // 不交换
				dp2[i] = common.Min(dp1[i-1], dp2[i-1]) + 1 // 交换
			} else {
				dp1[i] = dp1[i-1]     // 不交换，则上个位置也不能交换
				dp2[i] = dp2[i-1] + 1 // 交换，则上个位置也必须交换
			}
		} else {
			dp1[i] = dp2[i-1]     // 不交换，则上个位置必须交换
			dp2[i] = dp1[i-1] + 1 //  交换，则上个位置不能交换
		}
	}
	return common.Min(dp2[len(A)-1], dp1[len(B)-1])
}
