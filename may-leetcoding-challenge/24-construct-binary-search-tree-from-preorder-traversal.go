/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	idx := sort.Search(len(preorder), func(i int) bool { return preorder[i] > preorder[0] })

	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:idx]),
		Right: bstFromPreorder(preorder[idx:]),
	}
}
