package main

import "fmt"

func main() {
	a:= fib(9)
	fmt.Print(a)
}
//費波那契數列
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		next := a + b
		a, b = b, next
	}
	return a
}