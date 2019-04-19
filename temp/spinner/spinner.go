package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	res := []int{1, 1}
	for i := 2; i < x; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[x-1]
}

func fibslow(x int) int {
	if x < 2 {
		return x
	}
	return fibslow(x-1) + fibslow(x-2)
}
