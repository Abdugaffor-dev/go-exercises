package main

import "fmt"

func main() {
	FizzBuzz(15)
}

func FizzBuzz(n int) {
	if n > 0 {
		for i := 1; i <= n; i++ {
			if i%3 == 0 && i%5 != 0 {
				fmt.Println("Fizz")
			}
			if i%5 == 0 && i%3 != 0 {
				fmt.Println("Buzz")
			}
			if i%5 == 0 && i%3 == 0 {
				fmt.Println("FizzBuzz")
			}
		}
	}
}
