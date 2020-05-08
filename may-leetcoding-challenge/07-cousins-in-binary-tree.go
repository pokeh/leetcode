/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var depthX, depthY, parentX, parentY int

	dfs(root, x, 0, 0, &depthX, &parentX)
	dfs(root, y, 0, 0, &depthY, &parentY)

	return (depthX == depthY) && (parentX != parentY)
}

func dfs(node *TreeNode, target, depthTracker, parentValTracker int, depth, parentVal *int) {
	if node == nil {
		return
	}

	if node.Val == target {
		*depth = depthTracker
		*parentVal = parentValTracker
		return
	}

	dfs(node.Left, target, depthTracker+1, node.Val, depth, parentVal)
	dfs(node.Right, target, depthTracker+1, node.Val, depth, parentVal)
}
