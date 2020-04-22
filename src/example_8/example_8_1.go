package example_8

import (
	"fmt"
	"time"
)

func TestExample_8_1() {
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

func fib(n int) int {

	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
