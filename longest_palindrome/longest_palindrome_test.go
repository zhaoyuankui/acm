package main

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func Benchmark_longestPalindrome(b *testing.B) {
	s0 := randString(0)
	s1 := randString(1)
	s2 := randString(10)
	s3 := randString(100)
	s4 := randString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s0)
		longestPalindrome(s1)
		longestPalindrome(s2)
		longestPalindrome(s3)
		longestPalindrome(s4)
	}
}

func Benchmark_longestPalindrome2(b *testing.B) {
	s0 := randString(0)
	s1 := randString(1)
	s2 := randString(10)
	s3 := randString(100)
	s4 := randString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s0)
		longestPalindrome(s1)
		longestPalindrome(s2)
		longestPalindrome(s3)
		longestPalindrome(s4)
	}
}

const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randString(n int) string {
	if n == 0 {
		return ""
	}
	l := len(alpha)
	buff := bytes.Buffer{}
	for i := 0; i < n; i++ {
		buff.WriteRune(rune(alpha[rand.Intn(l)]))
	}
	return buff.String()
}

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}
