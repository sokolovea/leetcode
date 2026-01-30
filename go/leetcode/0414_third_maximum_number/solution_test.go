package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	inputSlice []int
	result     int
}{
	{[]int{}, 0},
	{[]int{5}, 5},
	{[]int{4, 5}, 5},
	{[]int{4, 5, 4}, 5},
	{[]int{4, 5, 3}, 3},
	{[]int{1, 5, -38, 2, -300, 6, -1000, 12, 41}, 6},
	{[]int{1, 5, -38, 2, -300, 6, -1000, 12, 41}, 6},
}

func TestThirdMaxDummy(t *testing.T) {
	for i, data := range testData {
		result := thirdMaxDummy(data.inputSlice)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: for %v result %v but should be %v", i, data.inputSlice, result, data.result))
	}
}

func TestThirdMaxOptimized(t *testing.T) {
	for i, data := range testData {
		result := thirdMaxOptimized(data.inputSlice)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: for %v result %v but should be %v", i, data.inputSlice, result, data.result))
	}
}

func internalBenchmarkThirdMax(sliceSize int) []int {
	inputSlice := make([]int, sliceSize)
	r := rand.New(rand.NewSource(42))
	for i := range inputSlice {
		inputSlice[i] = r.Int()
	}
	return inputSlice
}

var sizes = []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 4096, 16384, 65536, 262144, 1048576}

func BenchmarkThirdMaxDummy(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				inputSlice := internalBenchmarkThirdMax(size)
				b.StartTimer()
				_ = thirdMaxDummy(inputSlice)
			}
		})
	}
}

func BenchmarkThirdMaxOptimize(b *testing.B) {
	for _, size := range sizes {
		inputSlice := internalBenchmarkThirdMax(size)
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = thirdMaxOptimized(inputSlice)
			}
		})
	}
}
