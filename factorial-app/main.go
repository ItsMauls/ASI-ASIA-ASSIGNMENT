package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func powerOfTwo(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * 2
	}
	return result
}

func f(n int) int {
	fact := factorial(n)
	power := powerOfTwo(n)

	hasil := fact / power
	if fact%power != 0 {
		hasil = hasil + 1
	}

	return hasil
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("f(%d) = %d\n", i, f(i))
	}
}
