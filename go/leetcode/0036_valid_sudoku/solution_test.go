package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataSequence = []struct {
	inputSequence []byte
	result        bool
}{
	{[]byte{'1', '2'}, true},
	{[]byte{'1', '1'}, false},
	{[]byte{'1', '.', '.', '9'}, true},
	{[]byte{'1', '2', '.', '3', '.'}, true},
	{[]byte{'.', '.', '.', '.', '.'}, true},
	{[]byte{'.', '1', '.', '.', '1'}, false},
	{[]byte{'1', '.', '3', '2', '1'}, false},
	{[]byte{'9', '9', '9', '9', '.'}, false},
	{[]byte{}, true},
	{[]byte{'.'}, true},
}

var testDataSudoku = []struct {
	inputSudoku [][]byte
	result      bool
}{
	{[][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, true},
	{[][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, false},
}

func TestIsValidSequence(t *testing.T) {
	for i, data := range testDataSequence {
		result := isValidSequence(data.inputSequence)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: result for \"%v\" = %v but really is %v!", i, data.inputSequence, result, data.result))
	}
}

func TestIsValidSudoku(t *testing.T) {
	for i, data := range testDataSudoku {
		result := isValidSudoku(data.inputSudoku)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: result for \"%v\" = %v but really is %v!", i, data.inputSudoku, result, data.result))
	}
}

func BenchmarkSudokuFunctional(b *testing.B) {
	for _, data := range testDataSudoku {
		_ = isValidSudoku(data.inputSudoku)
	}
}
