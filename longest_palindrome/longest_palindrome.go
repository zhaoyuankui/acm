package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abcde"))
	fmt.Println(longestPalindrome("bbbbbbbbb"))
	fmt.Println(longestPalindrome("abaaabcbabaaa"))

	fmt.Println(longestPalindrome2(""))
	fmt.Println(longestPalindrome2("a"))
	fmt.Println(longestPalindrome2("abcde"))
	fmt.Println(longestPalindrome2("bbbbbbbbb"))
	fmt.Println(longestPalindrome2("abaaabcbabaaa"))
	return
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	max := 1
	maxIdx := 0
	flags := map[uint32]bool{}
	for i := 0; i < len(s); i++ {
		flags[makeKey(i, i)] = true
		flags[makeKey(i, i+1)] = true
	}
	for m := 2; m <= len(s); m++ {
		for n := 0; n < len(s); n++ {
			if n+m > len(s) {
				break
			}
			if s[n] == s[n+m-1] && flags[makeKey(n+1, n+m-1)] {
				flags[makeKey(n, n+m)] = true
				if max < m {
					max = m
					maxIdx = n
				}
			}
		}
	}
	return s[maxIdx : maxIdx+max]
}

func makeKey(i, j int) uint32 {
	return uint32(i)<<16 | uint32(j)
}

func longestPalindrome2(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	max := 1
	maxIdx := 0
	//	flags := make(map[uint32]bool, (l+1)*l/2)
	flags := map[uint32]bool{}
	for i := 0; i < l; i++ {
		flags[makeKey(i, i)] = true
		flags[makeKey(i, i+1)] = true
	}
	for m := 2; m <= l; m++ {
		for n := 0; n < l; n++ {
			if n+m > l {
				break
			}
			if _, ok := flags[makeKey(n, n+m)]; ok {
				continue
			}
			if s[n] == s[n+m-1] && flags[makeKey(n+1, n+m-1)] {
				flags[makeKey(n, n+m)] = true
				if max < m {
					max = m
					maxIdx = n
				}
			} else {
				r := n
				rr := l - (n + m)
				if rr > r {
					r = rr
				}
				for i := 1; i <= r; i++ {
					flags[makeKey(n-i, n+m+i)] = false
				}
			}
		}
	}
	return s[maxIdx : maxIdx+max]
}
