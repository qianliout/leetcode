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
	NIL string
	Sep string
}

func Constructor() Codec {
	return Codec{
		NIL: "null",
		Sep: ",",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return strings.Join(this.ser(root), this.Sep)
}

func (this *Codec) ser(root *TreeNode) []string {
	ans := make([]string, 0)
	if root == nil {
		ans = append(ans, this.NIL)
		return ans
	}
	ans = append(ans, fmt.Sprintf("%d", root.Val))
	left := this.ser(root.Left)
	right := this.ser(root.Right)
	ans = append(ans, left...)
	ans = append(ans, right...)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	da := strings.Split(data, this.Sep)

	return this.des(&da)
}

func (this *Codec) des(data *[]string) *TreeNode {
	if len(*data) == 0 {
		return nil
	}
	first := (*data)[0]
	*data = (*data)[1:]
	atoi, err := strconv.Atoi(first)
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: atoi}
	root.Left = this.des(data)
	root.Right = this.des(data)
	return root
}
