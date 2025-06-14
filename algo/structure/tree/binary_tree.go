package tree

func MaxDepth(root *TreeNode) int {
	var bfs func(*TreeNode, int)
	ans := 0
	bfs = func(root *TreeNode, high int) {
		if root == nil {
			return
		}

		if high > ans {
			ans = high
		}

		if root.Left != nil {
			bfs(root.Left, high+1)
		}
		if root.Right != nil {
			bfs(root.Right, high+1)
		}
	}

	bfs(root, 1)
	return ans
}
