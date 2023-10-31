package main

import "fmt"

func main() {
	fmt.Println(fib(2)) //1
	fmt.Println(fib(3)) //2
	fmt.Println(fib(4)) //3
}

// Just copied from internet
func fib(n int) int {
	//fl:=fibonacciLoop(n)
	return fibonacciRecursion(n)
}

func fibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}
