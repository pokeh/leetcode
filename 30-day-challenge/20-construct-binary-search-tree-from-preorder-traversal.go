func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := preorder[0]
	ptr := len(preorder)

	for i := range preorder {
		if root < preorder[i] {
			ptr = i
			break
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  bstFromPreorder(preorder[1:ptr]),
		Right: bstFromPreorder(preorder[ptr:]),
	}
}
