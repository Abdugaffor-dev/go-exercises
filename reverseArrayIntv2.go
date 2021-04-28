package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Printf("Before append: %d,%d\n", cap(arr), len(arr))
	fmt.Println(arr)

	arr = append(arr, 6, 7, 8)

	fmt.Printf("After append 1: %d,%d\n", cap(arr), len(arr))
	fmt.Println(arr)

	arr = reverseArray(arr)
	fmt.Printf("After reverse: %d,%d\n", cap(arr), len(arr))
	fmt.Println(arr)

	arr = reverseArray(arr)
	arr = append(arr, 9, 10)

	fmt.Printf("After append&reverse 2: %d,%d\n", cap(arr), len(arr))
	fmt.Println(arr)

}

func reverseArray(arr []int) []int {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}


//Results:

/*Before append: 5,5
[1 2 3 4 5]
After append 1: 10,8
[1 2 3 4 5 6 7 8]
After reverse: 10,8
[8 7 6 5 4 3 2 1]
After append&reverse 2: 10,10
[1 2 3 4 5 6 7 8 9 10]      */
