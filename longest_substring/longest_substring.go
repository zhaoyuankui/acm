package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aaa"))
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println(lengthOfLongestSubstring("abcbaaaabcdeabc"))
	fmt.Println(lengthOfLongestSubstring("abcbaaaabcdeabcdefghi"))
	fmt.Println(lengthOfLongestSubstring(""))
	return
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	flags := make(map[byte]int, 256)
	start, end, max := 0, 0, 0
	for ; end < len(s); end++ {
		c := byte(s[end])
		if idx, ok := flags[c]; !ok {
			flags[c] = end
		} else {
			flags[c] = end
			l := end - start
			if l > max {
				max = l
			}
			for i := start; i < idx; i++ {
				delete(flags, byte(s[i]))
			}
			start = idx + 1
		}
	}
	if end-start > max {
		return end - start
	}
	return max
}
