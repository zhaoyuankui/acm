package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyuankui/goutil/alg"
)

func Test_findMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 1.0, findMedianSortedArrays(alg.IntSlice(1), alg.IntSlice()))
	assert.Equal(t, 1.5, findMedianSortedArrays(alg.IntSlice(), alg.IntSlice(1, 2)))
	assert.Equal(t, 1.0, findMedianSortedArrays(alg.IntSlice(1), alg.IntSlice(1)))
	assert.Equal(t, 1.5, findMedianSortedArrays(alg.IntSlice(1), alg.IntSlice(2)))
	assert.Equal(t, 1.5, findMedianSortedArrays(alg.IntSlice(2), alg.IntSlice(1)))
	assert.Equal(t, 2.0, findMedianSortedArrays(alg.IntSlice(2, 3), alg.IntSlice(1)))
	assert.Equal(t, 2.0, findMedianSortedArrays(alg.IntSlice(1, 3), alg.IntSlice(2)))
	assert.Equal(t, 3.5, findMedianSortedArrays(alg.IntSlice(1, 3, 5), alg.IntSlice(2, 4, 6)))
	assert.Equal(t, 2.0, findMedianSortedArrays(alg.IntSlice(2), alg.IntSlice(1, 3)))
	assert.Equal(t, 2.5, findMedianSortedArrays(alg.IntSlice(1, 2), alg.IntSlice(3, 4)))
	assert.Equal(t, 6.5, findMedianSortedArrays(alg.IntSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11), alg.IntSlice(8)))
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(alg.IntSlice(1, 2, 3, 4, 5), alg.IntSlice(3, 4, 5, 6, 7))
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 || l2 == 0 {
		nums := nums1
		if l1 == 0 {
			nums = nums2
		}
		l := l1 + l2
		if l&1 > 0 {
			return float64(nums[l>>1])
		} else {
			return float64(nums[l>>1]+nums[l>>1-1]) / 2
		}
	}
	// Find the target in the longer sequence.
	if l2 > l1 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	half := (l1 + l2) >> 1
	p, q, i, j := 0, l1, l1>>1, 0
	//for i >= 0 && i <= l1 {
	for {
		j = half - i
		// Avoid exceed boundaries
		if j < 0 {
			q = i - 1
			i = (p + q) >> 1
			continue
		}
		if j > l2 {
			p = i + 1
			i = (p + q) >> 1
			continue
		}

		// Find the target.
		// The direction of i to move towards. false: left, true:right
		var d bool
		if i == 0 {
			if nums1[0] >= nums2[j-1] {
				break
			} else {
				d = true
			}
		} else if i == l1 {
			if nums1[l1-1] <= nums2[j] {
				break
			} else {
				d = false
			}
		} else if j == 0 {
			if nums2[0] >= nums1[i-1] {
				break
			} else {
				d = false
			}
		} else if j == l2 {
			if nums2[l2-1] <= nums1[i] {
				break
			} else {
				d = true
			}
		} else {
			if nums1[i] >= nums2[j-1] && nums1[i-1] <= nums2[j] {
				break
			} else if nums1[i] < nums2[j-1] {
				d = true
			} else {
				d = false
			}
		}

		if d {
			p = i + 1
			i = (p + q) >> 1
		} else {
			q = i - 1
			i = (p + q) >> 1
		}
	}
	// Construct result.
	if (l1+l2)&1 > 0 {
		if i == l1 {
			return float64(nums2[j])
		}
		if j == l2 {
			return float64(nums1[i])
		}
		if nums1[i] < nums2[j] {
			return float64(nums1[i])
		}
		return float64(nums2[j])
	}
	var mid1, mid2 int
	if i == 0 {
		mid1 = nums2[j-1]
	} else if j == 0 {
		mid1 = nums1[i-1]
	} else if nums1[i-1] > nums2[j-1] {
		mid1 = nums1[i-1]
	} else {
		mid1 = nums2[j-1]
	}
	if i == l1 {
		mid2 = nums2[j]
	} else if j == l2 {
		mid2 = nums1[i]
	} else if nums1[i] > nums2[j] {
		mid2 = nums2[j]
	} else {
		mid2 = nums1[i]
	}
	return float64(mid1+mid2) / 2
}
