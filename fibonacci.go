package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(6))
}
func Fibonacci(i int) int {
	if i == 1 || i == 2 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}
