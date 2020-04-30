/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := root.Val
	traverse(root, &maxSum)
	return maxSum
}

func traverse(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	lMax := traverse(node.Left, maxSum)
	rMax := traverse(node.Right, maxSum)

	*maxSum = max(
		*maxSum,
		node.Val,
		node.Val+lMax,
		node.Val+rMax,
		node.Val+lMax+rMax,
	)

	// choose the best path to pass on
	// (it'll be used to create a larger path, so you can't choose both left and right)
	return node.Val + max(lMax, rMax, 0)
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}
