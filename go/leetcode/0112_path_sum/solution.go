package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSumInternalDfs(root *TreeNode, currentSum int, targetSum int) bool {
	var result bool = false
	if root != nil {
		currentSum += root.Val
		if (root.Left == nil) && (root.Right == nil) {
			result = currentSum == targetSum
		} else {
			result = result || hasPathSumInternalDfs(root.Left, currentSum, targetSum)
			result = result || hasPathSumInternalDfs(root.Right, currentSum, targetSum)
		}
	}
	return result
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumInternalDfs(root, 0, targetSum)
}
