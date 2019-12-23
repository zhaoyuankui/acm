package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	assert.Equal(t, levelOrder(nil), [][]int{})
	node3 := &TreeNode{3, nil, nil}
	assert.Equal(t, levelOrder(node3), [][]int{[]int{3}})
	node2 := &TreeNode{2, nil, nil}
	node1 := &TreeNode{1, node2, node3}
	assert.Equal(t, levelOrder(node1), [][]int{[]int{1}, []int{2, 3}})
}

func Benchmark_levelOrder(b *testing.B) {
	// Initializations here.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrder(nil)
	}
}

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := &[][]int{}
	backTrace(res, root, 0)
	return *res
}

func backTrace(res *[][]int, node *TreeNode, d int) {
	if node == nil {
		return
	}
	if len(*res) < d+1 {
		*res = append(*res, []int{})
	}
	(*res)[d] = append((*res)[d], node.Val)
	backTrace(res, node.Left, d+1)
	backTrace(res, node.Right, d+1)
}
