package main

import (
	"fmt"
	"strconv"
	"strings"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	c := Constructor()
	data := c.serialize(nil)
	node := c.deserialize(data)
	fmt.Println(c.serialize(root))
	fmt.Println(node.Val)
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
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return c.NIL
	}
	ans := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no != nil {
				ans = append(ans, fmt.Sprintf("%d", no.Val))
				lev = append(lev, no.Left)
				lev = append(lev, no.Right)
			} else {
				ans = append(ans, c.NIL)
			}
		}
		queue = lev
	}

	return strings.Join(ans, c.Sep)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	des := strings.Split(data, c.Sep)
	if len(des) == 0 || des[0] == "" || des[0] == c.NIL {
		return nil
	}
	first := des[0]
	des = des[1:]

	a, _ := strconv.Atoi(first)
	root := &TreeNode{Val: a}
	stark := make([]*TreeNode, 0)
	stark = append(stark, root)
	for len(des) >= 2 && len(stark) > 0 {
		pre := stark[0]
		stark = stark[1:]
		fir, sec := des[0], des[1]
		des = des[2:]
		if fir != c.NIL {
			b, _ := strconv.Atoi(fir)
			left := &TreeNode{Val: b}
			pre.Left = left
			stark = append(stark, left)
		}
		if sec != c.NIL {
			b, _ := strconv.Atoi(sec)
			right := &TreeNode{Val: b}
			pre.Right = right
			stark = append(stark, right)
		}
	}

	return root
}
