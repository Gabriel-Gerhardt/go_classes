package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func inorderTraversal(root *TreeNode,) []int {
	var list []int
	iterateList(root, list)
	return list

}
func iterateList(root *TreeNode, list []int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	list = append(list, root.Val)
	inorderTraversal(root.Right)

	return list
}

}
