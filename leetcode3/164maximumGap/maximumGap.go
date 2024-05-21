package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	// fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	// fmt.Println(maximumGap([]int{1, 1, 1, 1}))
	fmt.Println(maximumGap([]int{1, 1, 1, 1, 1, 5, 5, 5, 5, 5}))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	mi, ma := nums[0], nums[0]
	for _, ch := range nums {
		mi = Min(mi, ch)
		ma = Max(ma, ch)
	}
	if ma == mi {
		return 0
	}

	// 单个桶的大小
	bucketSize := Max(1, (ma-mi)/(len(nums)-1))
	// 桶的个数
	bucketNum := (ma-mi)/bucketSize + 1

	bucket := make([]*entry, bucketNum)
	// 入桶
	for _, ch := range nums {
		idx := (ch - mi) / bucketSize
		PutToBucket(bucket, ch, idx)
	}
	// 遍历
	var pre *entry
	ans := 0
	for i := range bucket {
		en := bucket[i]
		if en != nil {
			ans = Max(ans, en.Ma-en.Mi)
			if pre != nil {
				ans = Max(ans, en.Mi-pre.Ma)
			}
			pre = en
		}
	}
	return ans
}

type entry struct {
	Ma, Mi int
}

func PutToBucket(bucket []*entry, val int, idx int) {
	en := bucket[idx]
	if en == nil {
		bucket[idx] = &entry{
			Ma: val,
			Mi: val,
		}
		return
	}
	if en.Mi > val {
		en.Mi = val
	}
	if en.Ma < val {
		en.Ma = val
	}
}
