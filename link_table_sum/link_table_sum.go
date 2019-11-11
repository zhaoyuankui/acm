package main

import (
	"bytes"
	"fmt"
)

func main() {
	l1 := makeList(9, 9, 9)
	l2 := makeList(9, 9, 9, 9, 9)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(addTwoNumbers(l1, l2))
	return
}

func makeList(nums ...int) *ListNode {
	var h *ListNode
	for i := len(nums) - 1; i >= 0; i -= 1 {
		node := &ListNode{
			Val:  nums[i],
			Next: h,
		}
		h = node
	}
	return h
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	var buf bytes.Buffer
	for {
		if l == nil {
			break
		}
		buf.WriteRune(rune(l.Val + '0'))
		l = l.Next
	}
	return buf.String()
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	promote := false
	var pre *ListNode
	p, q := l1, l2
	for ; p != nil && q != nil; p, q = p.Next, q.Next {
		pre = p
		sum := p.Val + q.Val
		if promote {
			sum += 1
		}
		if sum >= 10 {
			promote = true
			p.Val = sum - 10
		} else {
			promote = false
			p.Val = sum
		}
	}
	if p == nil {
		pre.Next = q
		p = q
	}
	for ; p != nil; p = p.Next {
		pre = p
		sum := p.Val
		if promote {
			sum += 1
		}
		if sum >= 10 {
			promote = true
			p.Val = sum - 10
		} else {
			promote = false
			p.Val = sum
		}
		if !promote {
			break
		}
	}
	if promote {
		pre.Next = &ListNode{
			Val: 1,
		}
	}
	return l1
}
