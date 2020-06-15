package algorithms

import (
	"testing"

	. "github.com/pokeh/leetcode/datastructs"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrderTraversal(t *testing.T) {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}

	assert.Equal(t, []int{1, 2, 3, 4, 5}, LevelOrderTraversal(&root))
}
