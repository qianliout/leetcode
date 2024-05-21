package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

/*
1:p 和 q 在 root 的子树中，且分列 root 的 异侧（即分别在左、右子树中）；
2:p=root，且 q 在 root 的左或右子树中；
3:q=root，且 p 在 root 的左或右子树中；
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//  这个怎么理解呢
	// 递归遍历左右子树，如果左右子树查到节点都不为空，则表明 p 和 q 分别在左右子树中，因此，当前节点即为最近公共祖先；
	if left != nil && right != nil {
		return root
	}
	if left != nil && right == nil {
		return left
	}
	if left == nil && right != nil {
		return right
	}
	return nil
}
