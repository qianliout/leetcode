package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func findKthNumber(n int, k int) int {
	p := 1 // 作为一个指针，指向当前所在位置，当p==k时，也就是到了排位第k的数
	prefix := 1
	for p < k {
		cnt := getCnt(prefix, k) // 获得当前前缀下所有子节点的和
		if p+cnt > k {           // 第k个数在当前前缀下
			prefix = prefix * 10
			p++ // 把指针指向了第一个子节点的位置，比如11乘10后变成110，指针从11指向了110
		} else if p+cnt <= k { // 第k个数不在当前前缀下
			prefix++
			p = p + cnt // 注意这里的操作，把指针指向了下一前缀的起点
		}
	}
	return prefix
}

// 该函数实现了统计范围 [1,limit] 内以 prefix 为前缀的数的个数。
func getCnt(prefix, limit int) int {
	nextPrefix := prefix + 1 // 下一个前缀
	curPrefix := prefix
	cnt := 0
	for curPrefix <= limit {
		// nextPrefix -curPrefix // 下一个前缀的起点减去当前前缀的起点
		// 当然，当 nextPrefix 的值大于上界的时候，那以这个前缀为根节点的十叉树就不是满十叉树了啊，应该到上界那里，后面都不再有子节点。
		// 因此，count+=nextPrefix−curPrefix 还是有些问题的，需要来修正这个问题:
		// 这里 limit-curPrefix+1中的加1是个难点，因为包括 limit 本身的值
		cnt = Min(nextPrefix-curPrefix, limit-curPrefix+1)
		// 如果说刚刚prefix是1，next是2，那么现在分别变成10和20
		// 1为前缀的子节点增加10个，十叉树增加一层, 变成了两层
		// 如果说现在prefix是10，next是20，那么现在分别变成100和200，
		// 1为前缀的子节点增加100个，十叉树又增加了一层，变成了三层
		curPrefix = curPrefix * 10
		nextPrefix = nextPrefix * 10
	}

	return cnt
}
