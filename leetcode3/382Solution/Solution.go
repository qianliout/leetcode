package main

import (
	"math/rand"
	"time"

	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

type Solution struct {
	Head *ListNode
	Data []*ListNode
	Rand *rand.Rand
}

func Constructor(head *ListNode) Solution {
	cur := head
	data := make([]*ListNode, 0)
	for cur != nil {
		data = append(data, cur)
		cur = cur.Next
	}

	return Solution{
		Data: data,
		Head: head,
		Rand: rand.New(rand.NewSource(time.Now().UnixMilli())),
	}
}

func (this *Solution) GetRandom() int {
	if this.Head == nil || len(this.Data) <= 0 {
		return 0
	}
	n := this.Rand.Intn(len(this.Data))
	return this.Data[n].Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
