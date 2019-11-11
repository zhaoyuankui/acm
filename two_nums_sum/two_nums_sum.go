package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum(list(6, 1, 4, 2, 5, 8, 3), 7))
	fmt.Println(twoSum(list(1, 4, 2, 3, 2, 5, 5, 12, 11), 10))
	fmt.Println(twoSum(list(1, 4, 2, 3, 2, 5, 5, 12, 11), 44))
	fmt.Println(twoSum(nil, 44))
	fmt.Println(twoSum(list(1), 44))
	fmt.Println(twoSum(list(), 44))
}

func list(nums ...int) []int {
	return nums
}

type sortNum struct {
	v   int
	idx int
}

type sortNumSlice []*sortNum

func (ss sortNumSlice) Less(i, j int) bool {
	return ss[i].v < ss[j].v
}

func (ss sortNumSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss sortNumSlice) Len() int {
	return len(ss)
}

func sortList(nums ...int) []*sortNum {
	sortedNums := make([]*sortNum, len(nums), len(nums))
	for i, n := range nums {
		sortedNums[i] = &sortNum{
			v:   n,
			idx: i,
		}
	}
	sort.Sort(sortNumSlice(sortedNums))
	return sortedNums
}

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}

	sortedNums := sortList(nums...)
	for i := 0; i < len(sortedNums); i += 1 {
		v := sortedNums[i].v
		t := target - v
		start, end := 0, len(sortedNums)-1
		for {
			mid := (start + end) / 2
			if sortedNums[mid].v > t {
				end = mid - 1
			} else if sortedNums[mid].v < t {
				start = mid + 1
			} else {
				if i == mid {
					start = start + 1
				} else {
					if sortedNums[i].idx < sortedNums[mid].idx {
						return []int{sortedNums[i].idx, sortedNums[mid].idx}
					} else {
						return []int{sortedNums[mid].idx, sortedNums[i].idx}
					}
				}
			}
			if start > end {
				break
			}
		}
	}
	return []int{-1, -1}
}
