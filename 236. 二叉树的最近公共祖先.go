package main

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 都没找到或者找到了p或者q，就返回
	if root == nil || root == p || root == q {
		// 直接返回root
		return root
	}
	// 在左右子树中找
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 注意是&&，左右子树都找到了
	if left != nil && right != nil {
		return root
	}
	// 只在左子树找到
	if left != nil {
		return left
	}
	// 只在右子树找到
	return right
}
