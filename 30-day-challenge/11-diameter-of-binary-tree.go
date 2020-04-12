func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := dfs(root, 0)
	return diameter
}

func dfs(node *TreeNode, currMax int) (depth int, nextMax int) {
	if node == nil {
		return
	}

	l, currMax := dfs(node.Left, currMax)
	r, currMax := dfs(node.Right, currMax)

	depth = max(l, r) + 1
	nextMax = max(l+r, currMax)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
