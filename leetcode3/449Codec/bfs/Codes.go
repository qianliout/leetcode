package main

import (
	"fmt"
	"strconv"
	"strings"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

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
// 相当于层序遍历
func (this *Codec) serialize(root *TreeNode) string {
	node := make([]*TreeNode, 0)
	node = append(node, root)
	ans := make([]string, 0)
	for len(node) > 0 {
		first := node[0]
		node = node[1:]
		if first != nil {
			ans = append(ans, fmt.Sprintf("%d", first.Val))
			node = append(node, first.Left)
			node = append(node, first.Right)
		} else {
			ans = append(ans, this.Null)
		}
	}
	return strings.Join(ans, this.sep)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data1 string) *TreeNode {
	data := strings.Split(data1, this.sep)
	if len(data) == 0 || data[0] == this.Null {
		return nil
	}
	first := data[0]
	data = data[1:]
	n, _ := strconv.Atoi(first)
	root := &TreeNode{Val: n}
	node := make([]*TreeNode, 0)
	node = append(node, root)
	for len(data) > 0 && len(node) > 0 {
		firstNode := node[0]
		node = node[1:]
		left := data[0]
		right := data[1]
		data = data[2:]

		lv, err := strconv.Atoi(left)
		if err == nil {
			firstNode.Left = &TreeNode{Val: lv}
			node = append(node, firstNode.Left)
		}
		rv, err := strconv.Atoi(right)
		if err == nil {
			firstNode.Right = &TreeNode{Val: rv}
			node = append(node, firstNode.Right)
		}
	}
	return root
}
