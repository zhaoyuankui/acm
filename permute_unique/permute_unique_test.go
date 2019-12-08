package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyuankui/goutil/alg"
)

func Test_permuteUnique(t *testing.T) {
	assert.Equal(t, [][]int{}, permuteUnique(alg.IntSlice()))
	assert.Equal(t, [][]int{[]int{0}}, permuteUnique(alg.IntSlice(0)))
	assert.Equal(t, 6, len(permuteUnique(alg.IntSlice(1, 2, 3))))
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 2, 3)), []int{1, 2, 3})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 2, 3)), []int{1, 3, 2})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 2, 3)), []int{2, 3, 1})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 2, 3)), []int{2, 1, 3})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 2, 3)), []int{3, 1, 2})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 2, 3)), []int{3, 2, 1})
	assert.Equal(t, 3, len(permuteUnique(alg.IntSlice(1, 1, 3))))
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 1, 3)), []int{1, 1, 3})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 1, 3)), []int{1, 3, 1})
	assert.Contains(t, permuteUnique(alg.IntSlice(1, 1, 3)), []int{3, 1, 1})
	assert.Equal(t, 6, len(permuteUnique(alg.IntSlice(2, 2, 1, 1))))
	assert.Contains(t, permuteUnique(alg.IntSlice(2, 2, 1, 1)), []int{2, 2, 1, 1})
	assert.Contains(t, permuteUnique(alg.IntSlice(2, 2, 1, 1)), []int{2, 1, 2, 1})
	assert.Contains(t, permuteUnique(alg.IntSlice(2, 2, 1, 1)), []int{1, 2, 1, 2})
	assert.Contains(t, permuteUnique(alg.IntSlice(2, 2, 1, 1)), []int{1, 1, 2, 2})
	assert.Contains(t, permuteUnique(alg.IntSlice(2, 2, 1, 1)), []int{1, 2, 2, 1})
	assert.Contains(t, permuteUnique(alg.IntSlice(2, 2, 1, 1)), []int{2, 1, 1, 2})
}

func Benchmark_permuteUnique(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	nums := []int{1, 1, 2, 3, 7, 5, 4, 3, 8, 6}
	for i := 0; i < b.N; i++ {
		permuteUnique(nums)
	}
}

type dualSlice [][]int

func (ds dualSlice) Len() int {
	return len(ds)
}

func (ds dualSlice) Less(i, j int) bool {
	for k := 0; k < len(ds[i]); k++ {
		if ds[i][k] > ds[j][k] {
			return false
		} else if ds[i][k] < ds[j][k] {
			return true
		}
	}
	return false
}

func (ds dualSlice) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds dualSlice) Unique() [][]int {
	res := [][]int{}
	if len(ds) == 0 {
		return res
	}
	res = append(res, ds[0])
	for i := 1; i < len(ds); i++ {
		for j := 0; j < len(ds[i]); j++ {
			if ds[i][j] != ds[i-1][j] {
				res = append(res, ds[i])
				break
			}
		}
	}
	return res
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	res = append(res, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		l := len(res)
		for j := 0; j < l; j++ {
			var k, p int
			resj := res[j]
			for k = 0; k < len(resj); k++ {
				if k > 0 {
					if p == nums[i] {
						p = resj[k]
						continue
					}
					res = append(res, append(append([]int{}, resj[:k]...), append([]int{nums[i]}, resj[k:]...)...))
				} else {
					res = append(res, append([]int{nums[i]}, resj...))
				}
				p = resj[k]
			}
			if p != nums[i] {
				res = append(res, append(resj, nums[i]))
			}
		}
		res = res[l:]
	}
	sort.Sort(dualSlice(res))
	return dualSlice(res).Unique()
}

func Test_permuteUnique2(t *testing.T) {
	/*
		assert.Equal(t, [][]int{}, permuteUnique2(alg.IntSlice()))
		assert.Equal(t, [][]int{[]int{0}}, permuteUnique2(alg.IntSlice(0)))
		assert.Equal(t, 6, len(permuteUnique2(alg.IntSlice(1, 2, 3))))
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 2, 3)), []int{1, 2, 3})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 2, 3)), []int{1, 3, 2})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 2, 3)), []int{2, 3, 1})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 2, 3)), []int{2, 1, 3})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 2, 3)), []int{3, 1, 2})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 2, 3)), []int{3, 2, 1})
		assert.Equal(t, 3, len(permuteUnique2(alg.IntSlice(1, 1, 3))))
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 1, 3)), []int{1, 1, 3})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 1, 3)), []int{1, 3, 1})
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 1, 3)), []int{3, 1, 1})
		assert.Equal(t, 6, len(permuteUnique2(alg.IntSlice(2, 2, 1, 1))))
		assert.Contains(t, permuteUnique2(alg.IntSlice(2, 2, 1, 1)), []int{2, 2, 1, 1})
		assert.Contains(t, permuteUnique2(alg.IntSlice(2, 2, 1, 1)), []int{2, 1, 2, 1})
		assert.Contains(t, permuteUnique2(alg.IntSlice(2, 2, 1, 1)), []int{1, 2, 1, 2})
		assert.Contains(t, permuteUnique2(alg.IntSlice(2, 2, 1, 1)), []int{1, 1, 2, 2})
		assert.Contains(t, permuteUnique2(alg.IntSlice(2, 2, 1, 1)), []int{1, 2, 2, 1})
		assert.Contains(t, permuteUnique2(alg.IntSlice(2, 2, 1, 1)), []int{2, 1, 1, 2})
		assert.Equal(t, 1, len(permuteUnique2(alg.IntSlice(1, 1, 1))))
		assert.Contains(t, permuteUnique2(alg.IntSlice(1, 1, 1)), []int{1, 1, 1})
	*/
	assert.Equal(t, 4, len(permuteUnique2(alg.IntSlice(1, 2, 2, 2))))
}

func Benchmark_permuteUnique2(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	nums := []int{1, 1, 2, 3, 7, 5, 4, 3, 8, 6}
	for i := 0; i < b.N; i++ {
		permuteUnique2(nums)
	}
}

type valueType [2]int

func permuteUnique2(nums []int) [][]valueType {
	res := [][]valueType{}
	if len(nums) == 0 {
		return res
	}
	flags := make(map[int]bool, len(nums))
	v := make([]valueType, len(nums))
	sort.Sort(sort.IntSlice(nums))
	backbrace(nums, flags, &res, v, 0)
	return res
}

func backbrace(nums []int, flags map[int]bool, res *[][]valueType, v []valueType, d int) {
	if d == len(nums) {
		*res = append(*res, append([]valueType{}, v...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if flags[i] {
			continue
		}
		if i != 0 && !flags[i-1] && nums[i] == nums[i-1] {
			continue
		}
		flags[i] = true
		v[d][0] = nums[i]
		v[d][1] = i
		backbrace(nums, flags, res, v, d+1)
		flags[i] = false
	}
}
