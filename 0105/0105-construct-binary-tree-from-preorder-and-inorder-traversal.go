package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, inorder)
}

// preorder: root -> left subtree -> right subtree
// inorder: left subtree -> root -> right subtree
func helper(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := preorder[0]

	if len(preorder) == 1 {
		// leaf node
		return &TreeNode{
			Val:   root,
			Left:  nil,
			Right: nil,
		}
	}

	inRootIdx := 0
	for i, v := range inorder {
		if v == root {
			inRootIdx = i
		}
	}

	// in preorder, right subtree starts after
	// root (length = 1) and left subtree (length = len(inorder[:inRootIdx]))
	preRightSubStartIdx := 1 + len(inorder[:inRootIdx])

	return &TreeNode{
		Val:   root,
		Left:  helper(preorder[1:preRightSubStartIdx], inorder[:inRootIdx]),
		Right: helper(preorder[preRightSubStartIdx:], inorder[inRootIdx+1:]),
	}
}
