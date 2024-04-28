package main

import (
	"fmt"
	"strconv"
	"strings"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	sep  string
	Null string
}

func Constructor() Codec {
	return Codec{
		sep:  ",",
		Null: "null",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.Null
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	ans := fmt.Sprintf("%d", root.Val) + this.sep + left + this.sep + right
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sl := strings.Split(data, this.sep)
	return des(&sl)
}

func des(data *[]string) *TreeNode {
	if len(*data) == 0 {
		return nil
	}

	first := (*data)[0]
	*data = (*data)[1:]

	n, err := strconv.Atoi(first)
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: n}
	root.Left = des(data)
	root.Right = des(data)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
