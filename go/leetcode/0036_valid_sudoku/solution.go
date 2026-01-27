package main

type BoardWalker func(board [][]byte, sequence []byte)

func isValidSequence(sequence []byte) bool {
	var numbers []byte = make([]byte, 10)
	var isValid bool = true
	for _, curCell := range sequence {
		if curCell < '0' || curCell > '9' {
			continue
		}
		curNumber := curCell - '0'
		if numbers[curNumber] > 0 {
			isValid = false
			break
		}
		numbers[curNumber] += 1
	}
	return isValid
}

func isValidSudoku(board [][]byte) bool {
	var isValid bool = true
	var sequence []byte = make([]byte, max(len(board), len(board[0])))
	for i := range len(board[0]) {
		func(board [][]byte, sequence []byte) {
			copy(sequence, board[i])
		}(board, sequence)
		isValid = isValidSequence(sequence[:len(board[0])])
		if !isValid {
			return false
		}
	}
	for j := range len(board) {
		func(board [][]byte, sequence []byte) {
			for elemIndex := range board[j] {
				sequence[elemIndex] = board[elemIndex][j]
			}
		}(board, sequence)
		isValid = isValidSequence(sequence[:len(board)])
		if !isValid {
			return false
		}
	}
	for i := range len(board) * len(board[0]) / (3 * 3) {
		squareCenterRow := (i/3)*3 + 1
		squareCenterCol := (i*3 + 1) % 9
		func(board [][]byte, sequence []byte) {
			counter := 0
			for i := squareCenterRow - 1; i <= squareCenterRow+1; i++ {
				for j := squareCenterCol - 1; j <= squareCenterCol+1; j++ {
					sequence[counter] = board[i][j]
					counter++
				}
			}
		}(board, sequence)
		isValid = isValidSequence(sequence[:len(board)])
		if !isValid {
			return false
		}
	}
	return isValid
}
