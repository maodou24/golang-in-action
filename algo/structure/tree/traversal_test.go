package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrderTraversal(t *testing.T) {
	s := PreOrderTraversal(
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 6},
			},
		},
	)

	assert.EqualValues(t, []int{3, 1, 5, 2, 6}, s)
}
