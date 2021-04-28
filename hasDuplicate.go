package main

import (
	"fmt"
)

func main() {
	var n = [6]int{22, 3, 12, 4, 2, 33}
	fmt.Println(hasDuplicate(n))
}

func hasDuplicate(n [6]int) string {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			if i != j {
				if n[i] == n[j] {
					return "There is duplicate"
				}
			}
		}
	}
	return "There is no duplicate"
}
