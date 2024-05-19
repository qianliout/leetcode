package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(smallestGoodBase("13"))
}

// 先不考虑二分
func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	// mMax := bits.Len(uint(nVal)) - 1 // bits.Len 用来获取表示目标数所需要的最小位数，输入3输出2（3 = 11）
	// 通过 bits.Len来限制枚举的深度
	ans := nVal - 1 // 最差的结果，也就是m进制表示结果为`11`的情况，此时 m进制的10 = n - 1；11 = n - 1 + 1 = n

	for m := 62; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		if k == 1 {
			// 如果等于1，那肯定能得到 n 的，但是1不会是结果
			continue
		}
		// 开m次方根
		mul, sum := 1, 1

		for i := 0; i < m; i++ {
			// 模拟 k^0 + k^1 + k^2 + ... + k^m-1
			mul *= k
			sum += mul
		}
		if sum == nVal {
			// 枚举完一个k，校验一遍是否刚好等于输入的数值
			// 这里可以直接返回 结果，但是是为什么没有想明白这里没有想明白
			// return strconv.Itoa(k)
			ans = min(ans, k)
		}
	}
	return strconv.Itoa(ans)
}
