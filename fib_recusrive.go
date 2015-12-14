package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number to get the fibonacci value")
	var i int
	fmt.Scanf("%d", &i)
	fmt.Println(fib(i))
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
