package tree

func PreOrderTraversal(root *TreeNode) (res []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
