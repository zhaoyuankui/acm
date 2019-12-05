package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyuankui/goutil/alg"
)

func Test_permute(t *testing.T) {
	assert.Equal(t, [][]int{}, permute(alg.IntSlice()))
	assert.Equal(t, [][]int{[]int{0}}, permute(alg.IntSlice(0)))
	assert.Equal(t, 6, len(permute(alg.IntSlice(1, 2, 3))))
	assert.Contains(t, permute(alg.IntSlice(1, 2, 3)), []int{1, 2, 3})
	assert.Contains(t, permute(alg.IntSlice(1, 2, 3)), []int{1, 3, 2})
	assert.Contains(t, permute(alg.IntSlice(1, 2, 3)), []int{2, 3, 1})
	assert.Contains(t, permute(alg.IntSlice(1, 2, 3)), []int{2, 1, 3})
	assert.Contains(t, permute(alg.IntSlice(1, 2, 3)), []int{3, 1, 2})
	assert.Contains(t, permute(alg.IntSlice(1, 2, 3)), []int{3, 2, 1})
}

func Benchmark_permute(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	nums := []int{1, 1, 2, 3, 7, 5, 4, 3, 8, 6}
	for i := 0; i < b.N; i++ {
		permute(nums)
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	res = append(res, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		l := len(res)
		for j := 0; j < l; j++ {
			var k int
			resj := res[j]
			for k = 0; k < len(resj); k++ {
				if k > 0 {
					res = append(res, append(append([]int{}, resj[:k]...), append([]int{nums[i]}, resj[k:]...)...))
				} else {
					res = append(res, append([]int{nums[i]}, resj...))
				}
			}
			res = append(res, append(resj, nums[i]))
		}
		res = res[l:]
	}
	return res
}
