package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	assert.Equal(t, findKthLargest([]int{1}, 1), 1)
	assert.Equal(t, findKthLargest([]int{1, 4, 7, 2, 5, 8}, 4), 4)
	assert.Equal(t, findKthLargest([]int{1, 4, 7, 2, 5, 8}, 1), 8)
	assert.Equal(t, findKthLargest([]int{1, 4, 7, 2, 5, 8}, 6), 1)
	/*
	 */
}

func Benchmark_findKthLargest(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest([]int{1, 4, 7, 2, 5, 8}, 4)
	}
}

func findKthLargest(nums []int, k int) int {
	partition(&nums, 0, len(nums)-1, k-1)
	return nums[k-1]
}

func partition(nums *[]int, l, r, k int) {
	p := (*nums)[l]
	l0, r0 := l, r
	for r > l {
		for ; r > l; r-- {
			if (*nums)[r] > p {
				(*nums)[l] = (*nums)[r]
				l++
				break
			}
		}
		for ; r > l; l++ {
			if (*nums)[l] <= p {
				(*nums)[r] = (*nums)[l]
				r--
				break
			}
		}
	}
	(*nums)[r] = p
	if k < r {
		partition(nums, l0, r-1, k)
	} else if k > r {
		partition(nums, r+1, r0, k)
	}
}

func Test_findKthLargest2(t *testing.T) {
	assert.Equal(t, findKthLargest2([]int{1}, 1), 1)
	assert.Equal(t, findKthLargest2([]int{1, 4, 7, 2, 5, 8}, 4), 4)
	assert.Equal(t, findKthLargest2([]int{1, 4, 7, 2, 5, 8}, 1), 8)
	assert.Equal(t, findKthLargest2([]int{1, 4, 7, 2, 5, 8}, 6), 1)
	/*
	 */
}

func Benchmark_findKthLargest2(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest2([]int{1, 4, 7, 2, 5, 8}, 4)
	}
}

func findKthLargest2(nums []int, k int) int {
	partition2(&nums, 0, len(nums)-1, k-1)
	return nums[k-1]
}

func partition2(nums *[]int, l, r, k int) {
	for {
		p := (*nums)[l]
		l0, r0 := l, r
		for r > l {
			for ; r > l; r-- {
				if (*nums)[r] > p {
					(*nums)[l] = (*nums)[r]
					l++
					break
				}
			}
			for ; r > l; l++ {
				if (*nums)[l] <= p {
					(*nums)[r] = (*nums)[l]
					r--
					break
				}
			}
		}
		(*nums)[r] = p
		if k == r {
			break
		}
		if k < r {
			l = l0
			r = r - 1
		} else if k > r {
			l = r + 1
			r = r0
		}
	}
}
