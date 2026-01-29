package main

func binarySearchSqrt(square int, left int, right int) int {
	center := (right + left + 1) / 2
	tempSquare := center * center
	switch {
	case tempSquare == square:
		return center
	case center == left || center == right:
		return left
	case tempSquare > square:
		return binarySearchSqrt(square, left, center)
	default:
		return binarySearchSqrt(square, center, right)
	}
}

func mySqrt(x int) int {
	return binarySearchSqrt(x, 0, x)
}

func isPerfectSquare(num int) bool {
	root := mySqrt(num)
	return root*root == num
}
