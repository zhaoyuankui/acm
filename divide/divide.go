package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(-7, 3))
	fmt.Println(divide(-7, 1))
	fmt.Println(divide(100, 4))
	fmt.Println(divide(int(-0x80000000), -1))
	fmt.Println(divide(2147483647, 1))
	fmt.Println(divide(1100540749, -1090366779))
	return
}

func divide(dividend, divisor int) int {
	// Overflow.
	if int32(dividend) == int32(-0x80000000) && int32(divisor) == -1 {
		return int(0x7fffffff)
	}
	dividend32 := int32(dividend)
	divisor32 := int32(divisor)
	minus := false
	if dividend32 < 0 {
		minus = !minus
	} else {
		dividend32 = -dividend32
	}
	if divisor32 < 0 {
		minus = !minus
	} else {
		divisor32 = -divisor32
	}
	res := getResult(dividend32, divisor32)
	if minus {
		return -res
	}
	return res
}

func getResult(dividend32, divisor32 int32) int {
	if dividend32 > divisor32 {
		return 0
	}
	res := 0
	mask := 0

redo:
	divisor := divisor32
	mask = 1
	dividend32 -= divisor
	for ; dividend32 <= 0 && divisor != -0x80000000; dividend32 -= divisor {
		res += mask
		mask = mask << 1
		divisor = divisor << 1
	}
	// recover the dividend32
	dividend32 += divisor
	if dividend32 <= divisor32 {
		goto redo
	}
	return res
}
