package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

type BSTIterator struct {
	Stark []*TreeNode
	Root  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{
		Stark: make([]*TreeNode, 0),
		Root:  root,
	}
	// 先把左子树全部入stark
	cur := root
	for cur != nil {
		bst.Stark = append(bst.Stark, cur)
		cur = cur.Left
	}

	return bst
}

func (this *BSTIterator) Next() int {
	if len(this.Stark) == 0 {
		return -1
	}

	last := this.Stark[len(this.Stark)-1]
	this.Stark = this.Stark[:len(this.Stark)-1]
	// 再把这个权的右树加入 stark中
	right := last.Right

	for right != nil {
		this.Stark = append(this.Stark, right)
		right = right.Left
	}
	return last.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stark) > 0
}
