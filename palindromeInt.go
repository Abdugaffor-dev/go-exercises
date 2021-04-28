package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(14541452541541))
}
func isPalindrome(num int) string {
	sum := 0
	r := 0
	for i := num; i > 0; i = i / 10 {
		r = i % 10
		sum = (sum * 10) + r
	}
	if num-sum == 0 {
		return "Is Palindrome!"
	}
	return "Is not Palindrome!"
}
