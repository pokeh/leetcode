func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxCount := 0

	for len(queue) > 0 {
		nodesInLevel := len(queue)
		count := 0
		isCounting := false

		for i := 0; i < nodesInLevel; i++ {
			// dequeue
			node := queue[0]
			queue = queue[1:]

			if !isCounting {
				// don't queue leading nulls
				if node == nil {
					continue
				}
				isCounting = true
			}
			count++

			if node == nil {
				queue = append(queue, nil, nil)
			} else {
				queue = append(queue, node.Left, node.Right)
			}
		}

		// remove trailing nulls
		for len(queue) > 0 && (queue[len(queue)-1] == nil) {
			queue = queue[:len(queue)-1]
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
