/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidSequence(root *TreeNode, arr []int) bool {
	return traverse(root, arr, 0)
}

func traverse(node *TreeNode, arr []int, pos int) bool {
	if node == nil || node.Val != arr[pos] {
		return false
	}

	if pos == len(arr)-1 {
		return node.Left == nil && node.Right == nil
	}

	return traverse(node.Left, arr, pos+1) || traverse(node.Right, arr, pos+1)
}
