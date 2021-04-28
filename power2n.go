package main

import (
	"fmt"
)

func main() {
	var n = 2
	fmt.Printf("%d th degree of 2 equals :%d\n", n, power2n(5))
}

func power2n(n int) int {
	p := 1
	for i := 1; i <= n; i++ {
		p = p * 2
	}
	return p
}
package main

import (
	"fmt"
)

func main() {
	var n = 2
	fmt.Printf("%d th degree of 2 equals :%d\n", n, power2n(5))
}

func power2n(n int) int {
	p := 1
	for i := 1; i <= n; i++ {
		p = p * 2
	}
	return p
}
