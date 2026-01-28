package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumInternalDfs(root *TreeNode, currentSum int, currentPath []int, allPaths *[][]int, targetSum int) {
	if root != nil {
		currentSum += root.Val
		currentPath = append(currentPath, root.Val)
		if (root.Left == nil) && (root.Right == nil) {
			if currentSum == targetSum {
				var newPath []int = make([]int, len(currentPath))
				copy(newPath, currentPath)
				*allPaths = append(*allPaths, newPath)
			}
		} else {
			pathSumInternalDfs(root.Left, currentSum, currentPath, allPaths, targetSum)
			pathSumInternalDfs(root.Right, currentSum, currentPath, allPaths, targetSum)
		}
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	allResultPaths := [][]int{}
	currentPath := []int{}
	if root == nil {
		return [][]int{}
	}
	pathSumInternalDfs(root, 0, currentPath, &allResultPaths, targetSum)
	return allResultPaths
}
