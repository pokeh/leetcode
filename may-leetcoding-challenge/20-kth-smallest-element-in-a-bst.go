/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	values := make([]int, 0)
	inorder(root, &values)
	return values[k-1]
}

func inorder(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, values)
	*values = append(*values, node.Val)
	inorder(node.Right, values)
}
