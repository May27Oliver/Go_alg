package main

import "fmt"

func main() {
	for i := 0 ; i < 11 ; i++ {
		fmt.Print(fibonacci(i)," ")
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}else if num == 1 {
		return 1
	}else{
		return fibonacci(num - 1)+ fibonacci(num-2)
	}
}