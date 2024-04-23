package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		le := make([]*Node, 0)
		res := make([]int, 0)
		for _, no := range queue {
			res = append(res, no.Val)
			for _, ch := range no.Children {
				if ch != nil {
					le = append(le, ch)
				}
			}
		}
		queue = le
		ans = append(ans, res)
	}
	return ans
}
