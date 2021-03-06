package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhaoyuankui/goutil/alg"
)

func main() {
	var a, b int
	for {
		if _, err := fmt.Scanf("%d %d\n", &a, &b); err != nil {
			break
		}
		fmt.Println(a + b)
	}
}

func Test_main(t *testing.T) {
	stdin, stdout := os.Stdin, os.Stdout
	os.Stdin, _ = os.Open("./test.in")
	os.Stdout, _ = os.Create("./test.out")

	main()

	// Check
	alg.AssertEqual(t, "./test.out", []string{"3", "7"})
	alg.AssertEqual(t, "./test.out", []int{3, 7})

	os.Stdin, os.Stdout = stdin, stdout
}

func Benchmark_main(b *testing.B) {
	// Initializations here.
	stdin, stdout := os.Stdin, os.Stdout
	os.Stdin, _ = os.Open("./test.in")
	os.Stdout, _ = os.Open("/dev/null")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		os.Stdin.Seek(0, 0)
		main()
	}

	os.Stdin, os.Stdout = stdin, stdout
}
